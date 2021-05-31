package main

import (
	"github.com/gorilla/websocket"
	_ "github.com/gorilla/websocket"
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

func newClient(socket *websocket.Conn, id string) chan client {
	newChannelClient := make(chan client)
	go clientChanSign(newChannelClient, socket, id)
	return newChannelClient
}

func clientChanSign(newChannelClient chan client, socket *websocket.Conn, id string) {
	newChannelClient <- client{
		Id:     id,
		Socket: socket,
		Sender: false,
	}
}
