package main

import (
	"encoding/json"
	"log"
	"time"

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
	Message Message
}

func (client *client) ReadMessagePool() {
	for {
		time.Sleep(tick * time.Second)
		log.Println("lendo")
		go func() {
			msgType, msg, _ := client.Socket.ReadMessage()
			client.Sender <- true
			_ = json.Unmarshal(msg, &client.Message)
			client.Message.MsgType = msgType
			log.Println(client.Message)
		}()
	}
}
func (client *client) WriteMessagePool() {
	for {
		time.Sleep(tick * time.Second)
		select {
		case sender := <-client.Sender:
			log.Println("opaa")
			if sender {
				go func() {

					var msg Message = client.Message
					log.Println(hub.Registered[string(msg.IdReciever)])
					for key, value := range hub.Registered {
						if key == msg.IdReciever {
							log.Print("Matchhhh: ")
							log.Println("Key: ", key, ", Value: ", value)
							value.Socket.WriteMessage(msg.MsgType, []byte(msg.Message))
						}

					}
				}()
			}
		}
	}
}

func newClient(socket *websocket.Conn, id string) client {
	var senderChan chan bool = make(chan bool)
	go func() { senderChan <- false }()
	client := client{
		Id:     id,
		Socket: socket,
		Sender: senderChan,
	}
	return client
}

func clientChanSign(newChannelClient chan client, socket *websocket.Conn, id string) client {
	var senderChan chan bool = make(chan bool)
	senderChan <- false
	return client{
		Id:     id,
		Socket: socket,
		Sender: senderChan,
	}
}
