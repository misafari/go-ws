package main

import (
	"log"
	"net/http"
	"ws/internal/handlers"
)

func main() {
	routes := routes()

	log.Println("Starting listen to websocket connections")
	go handlers.WsIncomeMessageChannelListener()

	log.Println("Starting Web server on port 8080")
	_ = http.ListenAndServe(":8080", routes)
}
