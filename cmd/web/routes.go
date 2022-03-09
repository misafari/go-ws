package main

import (
	"github.com/bmizerany/pat"
	"net/http"
	"ws/internal/handlers"
)

func routes() http.Handler {
	mux := pat.New()
	mux.Get("/", http.HandlerFunc(handlers.Home))
	mux.Get("/ws", http.HandlerFunc(handlers.WsEndpoint))

	fileServer := http.FileServer(http.Dir("./html/assets/"))
	mux.Get("/assets/", http.StripPrefix("/assets", fileServer))

	return mux
}
