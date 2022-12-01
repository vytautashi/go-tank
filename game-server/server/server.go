package server

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
	"github.com/vytautashi/go-tank/game-server/client"
	"github.com/vytautashi/go-tank/game-server/util"
)

const (
	framesPerSecond            = 30
	frameDurationInMiliseconds = 1000 / framesPerSecond
)

type Server struct {
	game       *Game
	maxPlayers int
	nowPlayers int

	// Registered clients in server
	clients map[uint32]*client.Client

	// Receiving messages/commands to this channel from clients
	commandsFromClients chan util.ChannelMessage

	// Receiving command from clients for removal from server
	deleteClients chan uint32

	// Receiving requests for registering clients
	addClients chan *websocket.Conn
}

func New(maxPlayers int) *Server {
	s := &Server{
		game:                NewGame(600, 600),
		maxPlayers:          maxPlayers,
		nowPlayers:          0,
		clients:             make(map[uint32]*client.Client, 600),
		commandsFromClients: make(chan util.ChannelMessage, 1000),
		deleteClients:       make(chan uint32, 32),
		addClients:          make(chan *websocket.Conn, 32),
	}

	go s.run()
	return s
}

// Used by controller for putting request to register client
func (s *Server) ClientRegister(conn *websocket.Conn) {
	s.addClients <- conn
}

// Main proccess/loop of the server
func (s *Server) run() {
	ticker := time.NewTicker(time.Duration(frameDurationInMiliseconds) * time.Millisecond)

	for {
		select {
		case <-ticker.C: // Every frame updates game
			s.removeDisconnectedClients()
			s.registerNewClients()
			s.game.update()
			s.sendInfoToClients()
			s.sendBulletsPositionsToClients()

		case msg := <-s.commandsFromClients:
			clientID := msg.Sender
			player := s.game.players[clientID]
			commandType := msg.Data[0]
			commandData := msg.Data[1:]

			switch commandType {
			case CommandFromClientUpdateInput:
				player.updateInput(commandData[0])

			default:
				log.Printf("received unknown command from player(id: %d)\n", clientID)

			}
		}
	}
}

func (s *Server) removeDisconnectedClients() {
	for {
		select {
		case clientID := <-s.deleteClients:
			// Removes client from server and game
			delete(s.clients, clientID)
			delete(s.game.players, clientID)

		default:
			s.nowPlayers = len(s.clients)
			return
		}
	}
}

func (s *Server) registerNewClients() {
	for {
		select {
		case conn := <-s.addClients:
			// Stops registering new clients, when exceeds limit
			if s.nowPlayers >= s.maxPlayers {
				conn.Close()
				continue
			}
			s.nowPlayers++

			// Creates client
			id := uint32(nextClientID())
			c := client.New(conn, s.commandsFromClients, s.deleteClients, id)
			s.clients[id] = c
			c.Run()

			// Creates player
			s.game.players[id] = NewPlayer()

			// Send initial data to client
			msg := util.ToBytes(CommandToClientInitPlayer, id)
			c.Send(msg)
			msg2 := util.ToBytes(CommandToClientInitMap, s.game.mapWidth, s.game.mapheight)
			c.Send(msg2)

		default:
			s.nowPlayers = len(s.clients)
			return
		}
	}
}

// Send to clients information about all player position in game
func (s *Server) sendInfoToClients() {
	msg := util.ToBytes(CommandToClientPlayersPositions)

	for id, player := range s.game.players {
		msg = util.ToBytes(msg, id, player.x, player.y)
	}

	for _, c := range s.clients {
		c.Send(msg)
	}
}

// Send to clients information about all bullets positions in game
func (s *Server) sendBulletsPositionsToClients() {
	msg := util.ToBytes(
		CommandToClientBulletsPositions,
		s.game.bulletManager.getBulletsData())

	for _, c := range s.clients {
		c.Send(msg)
	}
}
