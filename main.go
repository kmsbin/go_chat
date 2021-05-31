package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/websocket"
	_ "github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

var clients = make(map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan Message)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func main() {
	newHub := newHub()
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	go newHub.Run()

	r.Get("/enter-room", func(w http.ResponseWriter, r *http.Request) {
		uuidToken := uuid.NewV4()
		w.Write([]byte(uuidToken.String()))

		http.Redirect(w, r, "localhost:8080/ws", http.StatusMovedPermanently)
	})
	r.Get("/ws/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "id")
		log.Println(id)
		socket, _ := upgrader.Upgrade(w, r, nil)
		newHub.Register(newClient(socket, id))
	})

	http.ListenAndServe(":8080", r)
}
