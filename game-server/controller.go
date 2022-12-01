package main

import (
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/vytautashi/go-tank/game-server/server"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func getWS(s *server.Server) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Upgrades to websocket connection
		var conn, err = upgrader.Upgrade(w, r, nil)
		if err != nil {
			return
		}
		s.ClientRegister(conn)
	}
}
