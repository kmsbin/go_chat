package main

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
)

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
		newMsg.IdSender = client.Id
		if msgType == websocket.TextMessage {
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
func newClient(socket *websocket.Conn, id string) Client {
	var senderChan chan bool = make(chan bool, 10)
	go func() { senderChan <- false }()
	client := Client{
		Id:      id,
		Socket:  socket,
		Sender:  senderChan,
		Message: make(chan Message),
	}
	return client
}

// func clientChanSign(newChannelClient chan client, socket *websocket.Conn, id string) client {
// 	var senderChan chan bool = make(chan bool)
// 	senderChan <- false
// 	return client{
// 		Id:     id,
// 		Socket: socket,
// 		Sender: senderChan,
// 	}
// }
