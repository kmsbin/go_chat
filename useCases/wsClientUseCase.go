package useCases

import (
	"encoding/json"
	"go_chat/entities"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	Id       string          `json:"id"`
	Username string          `json:"username"`
	Socket   *websocket.Conn `json:"_"`
	Sender   chan bool
	Message  chan entities.Message
	hub      *Hub
}

func (client *Client) ReadMessagePool() {
	for {
		log.Println("lendo")
		msgType, msg, err := client.Socket.ReadMessage()

		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			client.hub.Unregister(*client)
			break
		}

		var newMsg entities.Message
		_ = json.Unmarshal(msg, &newMsg)
		if msgType == websocket.TextMessage {
			newMsg.Username = client.Username
			newMsg.IdSender = client.Id
			client.Message <- newMsg
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
		for key, value := range client.hub.Registered {
			if key == msg.IdReciever {
				parsedMsg, err := json.Marshal(msg)
				if err == nil {
					value.Socket.WriteMessage(websocket.TextMessage, parsedMsg)
				}
			}

		}
	}
}
func NewClient(socket *websocket.Conn, id string, username string, hub *Hub) Client {
	var senderChan chan bool = make(chan bool, 10)
	return Client{
		Id:       id,
		Socket:   socket,
		Sender:   senderChan,
		Message:  make(chan entities.Message),
		Username: username,
		hub:      hub,
	}
}
