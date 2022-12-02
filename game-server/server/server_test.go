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
