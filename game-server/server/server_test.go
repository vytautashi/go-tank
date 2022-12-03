package server

import (
	"testing"

	"github.com/gorilla/websocket"
)

// Test: Constants
func TestConstantsServer(t *testing.T) {
	// 1: framesPerSecond
	expect := 30
	result := framesPerSecond
	if result != expect {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}

	// 2: frameDurationInMiliseconds
	expect = 33
	result = frameDurationInMiliseconds
	if result != expect {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}

	// 3: frameDurationInMiliseconds = 1000 / framesPerSecond
	expect = 1000 / framesPerSecond
	result = frameDurationInMiliseconds
	if result != expect {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}
}

// Test: New()
func TestNewServer(t *testing.T) {
	srv := New(100)
	expect := 100
	result := srv.maxPlayers
	if result != expect {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}
	expect = 0
	result = srv.nowPlayers
	if result != expect {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}
}

// Test: ClientRegister()
func TestClientRegister(t *testing.T) {
	srv := New(100)
	data := websocket.Conn{}
	go srv.ClientRegister(&data)

	expect := &data
	result := <-srv.addClients
	if result != expect {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}
}

// Test: removeDisconnectedClients()
func TestRemoveDisconnectedClients(t *testing.T) {
	// 0: Initial setup
	s := New(100)
	// Player 1
	s.clients[1] = nil
	s.game.players[1] = nil
	// Player 2
	s.clients[2] = nil
	s.game.players[2] = nil
	// Player 3
	s.clients[3] = nil
	s.game.players[3] = nil
	s.nowPlayers = 3

	expect := 3
	result := len(s.clients)
	if result != expect {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}
	expect = 3
	result = len(s.game.players)
	if result != expect {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}
	expect2 := true
	_, result2 := s.clients[2]
	if result2 != expect2 {
		t.Fatalf(`result2 = %v, expect2 = %v`, result2, expect2)
	}
	expect2 = true
	_, result2 = s.game.players[2]
	if result2 != expect2 {
		t.Fatalf(`result2 = %v, expect2 = %v`, result2, expect2)
	}

	// 1: Deletes player/client with id:2
	s.deleteClients <- 2
	s.removeDisconnectedClients()
	expect = 2
	result = len(s.clients)
	if result != expect {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}
	expect = 2
	result = len(s.game.players)
	if result != expect {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}
	expect = 2
	result = s.nowPlayers
	if result != expect {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}
	expect2 = false
	_, result2 = s.clients[2]
	if result2 != expect2 {
		t.Fatalf(`result2 = %v, expect2 = %v`, result2, expect2)
	}
	expect2 = false
	_, result2 = s.game.players[2]
	if result2 != expect2 {
		t.Fatalf(`result2 = %v, expect2 = %v`, result2, expect2)
	}

	// 2: Deletes 2 players/clients
	s.deleteClients <- 1
	s.deleteClients <- 3
	s.removeDisconnectedClients()
	expect = 0
	result = len(s.clients)
	if result != expect {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}
	expect = 0
	result = len(s.game.players)
	if result != expect {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}
	expect = 0
	result = s.nowPlayers
	if result != expect {
		t.Fatalf(`result = %v, expect = %v`, result, expect)
	}
}
