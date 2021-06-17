package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
)

var hub Hub = newHub()

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func serveWs(hub *Hub, w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	username := chi.URLParam(r, "username")
	log.Println(username)
	socket, _ := upgrader.Upgrade(w, r, nil)
	client := newClient(socket, id, username)
	log.Println(id)

	hub.Register(client)
}

type Client struct {
	Id       string          `json:"id"`
	Username string          `json:"username"`
	Socket   *websocket.Conn `json:"_"`
	Sender   chan bool       `json:"_"`
	Message  chan Message    `json:"_"`
}

func (client *Client) ReadMessagePool() {
	for {
		log.Println("lendo")
		msgType, msg, err := client.Socket.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		var newMsg Message
		_ = json.Unmarshal(msg, &newMsg)
		if msgType == websocket.TextMessage {
			newMsg.Username = client.Username
			newMsg.IdSender = client.Id
			client.Message <- newMsg
			log.Println(client.Message)
			log.Println("New message: ", newMsg)
		}

	}
}
func (client *Client) WriteMessagePool() {
	for {
		msg, err := <-client.Message
		if !err {
			log.Println("error")
		}
		log.Println("opaa")
		for key, value := range hub.Registered {
			if key == msg.IdReciever {
				parsedMsg, err := json.Marshal(msg)
				if err == nil {
					value.Socket.WriteMessage(websocket.TextMessage, parsedMsg)
				}
			}

		}
	}
}
func newClient(socket *websocket.Conn, id string, username string) Client {
	var senderChan chan bool = make(chan bool, 10)
	return Client{
		Id:       id,
		Socket:   socket,
		Sender:   senderChan,
		Message:  make(chan Message),
		Username: username,
	}
}
