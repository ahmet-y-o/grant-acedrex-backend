package main

import (
	"acedrex/ws"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func main() {

	ws.InitRoomService()

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		// AllowedOrigins:   []string{"https://foo.com"}, // Use this to allow specific origin hosts
		// AllowedOrigins: []string{"https://*", "http://*"},
		AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))

	// Register the handler function for the root URL
	r.HandleFunc("/", ws.GetRoomsList)
	r.Get("/create-room", ws.CreateRoom)
	r.HandleFunc("/join-room/{roomId}", ws.JoinRoom)
	r.HandleFunc("/isFull/{roomId}", ws.IsRoomFull)

	// Start the server on port 8080
	http.ListenAndServe(":8080", r)

}
