package client

import (
	"time"

	"github.com/gorilla/websocket"
	"github.com/vytautashi/go-tank/game-server/util"
)

const (
	msgLengthMin = 2
	msgLengthMax = 8
	timeoutLimit = 5 * time.Second
	pingingTime  = 1 * time.Second
)

type Client struct {
	id         uint32
	conn       *websocket.Conn
	serverChan chan<- util.ChannelMessage
	deleteChan chan<- uint32
	send       chan []byte
}

// Creates new `Client` and returns it.
func New(
	conn *websocket.Conn,
	serverChan chan<- util.ChannelMessage,
	deleteChan chan<- uint32,
	id uint32) *Client {

	return &Client{
		id:         id,
		conn:       conn,
		serverChan: serverChan,
		deleteChan: deleteChan,
		send:       make(chan []byte, 16),
	}
}

// Create go routines for client for receiving and
// sending messages via websockets
func (c *Client) Run() {
	go c.read()
	go c.write()
}

// Function to send `message` to client using `send` channel
func (c *Client) Send(msg []byte) {
	util.SendBytesToChannel(c.send, msg)
}

// Used for sending `messages` to clients via websocket
// that were received via channel `send`.
func (c *Client) write() {
	ticker := time.NewTicker(pingingTime)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()

	for {
		select {
		case msg := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(timeoutLimit))
			err := c.conn.WriteMessage(websocket.BinaryMessage, msg)
			if err != nil {
				return
			}

		case <-ticker.C: // Pinging to check if connection still alive
			c.conn.SetWriteDeadline(time.Now().Add(timeoutLimit))
			err := c.conn.WriteMessage(websocket.PingMessage, nil)
			if err != nil {
				return
			}
		}
	}
}

// Used for reading incoming `messages` from client websocket
func (c *Client) read() {
	defer func() {
		c.conn.Close()
		c.deleteChan <- c.id
	}()
	c.conn.SetReadLimit(msgLengthMax)
	c.conn.SetReadDeadline(time.Now().Add(timeoutLimit))
	c.conn.SetPongHandler(c.pongHandler)

	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			return
		}

		// If message too big or too small ignore `message`
		if len(msg) > msgLengthMax || len(msg) < msgLengthMin {
			continue
		}

		// Sends received `message` to server
		util.SendToChannelMessage(c.serverChan, c.id, msg)
	}
}

func (c *Client) pongHandler(appData string) error {
	c.conn.SetReadDeadline(time.Now().Add(timeoutLimit))
	return nil
}
