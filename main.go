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

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/ws-handshake", func(w http.ResponseWriter, r *http.Request) {
		socket, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println(err)
		}
		for {
			// Vamos ler a mensagem recebida via Websocket
			msgType, msg, err := socket.ReadMessage()
			if err != nil {
				fmt.Println(err)
				return
			}
			log.Println(msgType)
			log.Println("Mensagem recebida: ", string(msg))
			err = socket.WriteMessage(msgType, msg)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	})

	http.ListenAndServe(":8080", r)
}
