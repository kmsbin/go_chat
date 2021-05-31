package main

import (
	"log"
)

type Hub struct {
	Registered   map[string]client
	Unregistered map[string]client
	History      map[string]string
}

func newHub() *Hub {
	return &Hub{
		Registered:   make(map[string]client, 10),
		Unregistered: make(map[string]client, 10),
		History:      make(map[string]string, 10),
	}
}

func (Hub *Hub) Register(newClient chan client) string {
	syncClient, _ := <-newClient
	Hub.Registered[syncClient.Id] = syncClient
	return syncClient.Id
}
func (Hub *Hub) Unregister(clientChan chan client) {
	newClient := <-clientChan

	delete(Hub.Registered, newClient.Id)
	Hub.Unregistered[newClient.Id] = newClient
}
func (Hub *Hub) Run() {
	for {
		for _, newClient := range Hub.Registered {

			_, msg, err := newClient.Socket.ReadMessage()

			if err != nil {
				log.Println(err)
				return
			}

			log.Println("Mensagem recebida: ", string(msg))
			// logger := []byte{byte(len(Hub.Registered))}
			// log.Println(logger)
			// newClient.Socket.WriteMessage(msgType, msg)

			// log.Println(len(Hub.Registered))
		}
	}
}
