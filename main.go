package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/gorilla/websocket"
	_ "github.com/gorilla/websocket"
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
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/ws", func(w http.ResponseWriter, r *http.Request) {
		socket, err := upgrader.Upgrade(w, r, nil)
		clients[socket] = true
		if err != nil {
			fmt.Println(err)
		}
		for {
			// Vamos ler a mensagem recebida via Websocket
			msgType, msg, err := socket.ReadMessage()

			for client, _ := range clients {

				if err != nil {
					fmt.Println(err)
					return
				}
				log.Println("Mensagem recebida: ", string(msg))
				logger := []byte{byte(len(clients))}
				log.Println(logger)

				err = client.WriteMessage(msgType, logger)
				err = client.WriteMessage(msgType, msg)
				if err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	})

	http.ListenAndServe(":8080", r)
}
