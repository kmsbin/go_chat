package main

import (
	"encoding/json"
	"log"

	"github.com/gorilla/websocket"
	_ "github.com/gorilla/websocket"
)

const (
	tick = 1
)

type event struct {
	Message  []byte
	Register client
}

type client struct {
	Id      string
	Socket  *websocket.Conn
	Sender  chan bool
	Message chan Message
}

func (client *client) ReadMessagePool() {
	for {
		// time.Sleep(tick * time.Second)
		// go func() {
		log.Println("lendo")
		msgType, msg, _ := client.Socket.ReadMessage()

		var newMsg Message
		_ = json.Unmarshal(msg, &newMsg)
		newMsg.MsgType = msgType
		newMsg.IdSender = client.Id
		if msgType > 0 {
			client.Message <- newMsg
			log.Println(client.Message)
			log.Println("New message: ", newMsg)
		}

		// }()
	}
}
func (client *client) WriteMessagePool() {
	for {
		// time.Sleep(tick * time.Second)
		msg := <-client.Message
		log.Println("opaa")
		// if sender {
		// log.Println(hub.Registered[string(msg.IdReciever)])
		for key, value := range hub.Registered {
			if key == msg.IdReciever {
				log.Println("Key: ", key, ", Value: ", value)
				parsedMsg, err := json.Marshal(msg)
				if err == nil {
					value.Socket.WriteMessage(msg.MsgType, parsedMsg)
				}
			}

		}
		// }
	}
}
func newClient(socket *websocket.Conn, id string) client {
	var senderChan chan bool = make(chan bool, 10)
	go func() { senderChan <- false }()
	client := client{
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
