package main

import (
	"log"
	"net/http"

	"github.com/vytautashi/go-tank/game-server/server"
)

func main() {
	s := server.New(200)

	http.HandleFunc("/ws", getWS(s))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
