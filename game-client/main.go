package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("host: localhost:8081")

	http.Handle("/", http.FileServer(http.Dir(".")))

	err := http.ListenAndServe(":8081", nil)
	if err != nil {
		log.Fatal(err)
	}
}
