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
