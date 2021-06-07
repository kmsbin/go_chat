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
var hub Hub = newHub()

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	go hub.Run()

	r.Get("/enter-room", func(w http.ResponseWriter, r *http.Request) {
		uuidToken := uuid.NewV4()
		w.Write([]byte(uuidToken.String()))

		http.Redirect(w, r, "localhost:8080/ws", http.StatusMovedPermanently)
	})
	r.Get("/ws/{id}", func(w http.ResponseWriter, r *http.Request) {
		serveWs(&hub, w, r)
	})

	http.ListenAndServe(":8080", r)
}
func serveWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	socket, _ := upgrader.Upgrade(w, r, nil)
	client := newClient(socket, id)
	go client.ReadMessagePool()
	go client.WriteMessagePool()
	log.Println(id)

	hub.Register(client)
}
