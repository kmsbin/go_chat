package main

import (
	"github.com/gorilla/websocket"
	_ "github.com/gorilla/websocket"
	uuid "github.com/satori/go.uuid"
)

type event struct {
	Message  []byte
	Register client
}

type client struct {
	Id     string
	Socket *websocket.Conn
	Sender bool
}

func newClient(socket *websocket.Conn) chan client {
	newChannelClient := make(chan client)
	go clientChanSign(newChannelClient, socket)
	return newChannelClient
}

func clientChanSign(newChannelClient chan client, socket *websocket.Conn) {
	newChannelClient <- client{
		Id:     uuid.NewV4().String(),
		Socket: socket,
		Sender: false,
	}
}
