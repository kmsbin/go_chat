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

type Message struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

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
	})
	r.Get("/ws", func(w http.ResponseWriter, r *http.Request) {
		socket, _ := upgrader.Upgrade(w, r, nil)
		log.Println("ASFDASDF")
		clientId := newHub.Register(newClient(socket))
		w.Write([]byte(clientId))
	})

	http.ListenAndServe(":8080", r)
}
