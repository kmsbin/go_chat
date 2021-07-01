package useCases

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

type Hub struct {
	Registered   map[string]Client `json:"registered-user,omitempty"`
	Unregistered map[string]Client `json:"unregistered-user,omitempty"`
	History      map[string]string `json:"-"`
}

func NewHub() Hub {
	return Hub{
		Registered:   make(map[string]Client, 10),
		Unregistered: make(map[string]Client, 10),
		History:      make(map[string]string, 10),
	}
}

func (Hub *Hub) Register(syncClient Client) string {
	go syncClient.ReadMessagePool()
	go syncClient.WriteMessagePool()
	Hub.Registered[syncClient.Id] = syncClient
	return syncClient.Id
}
func (Hub *Hub) Unregister(NewClient Client) {

	for key := range Hub.Registered {
		if key == NewClient.Id {
			delete(Hub.Registered, NewClient.Id)
			Hub.Unregistered[NewClient.Id] = NewClient
		}
	}
}
func (Hub *Hub) Run() {
}
func (hub *Hub) ServeWs(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	username := chi.URLParam(r, "username")
	socket, _ := upgrader.Upgrade(w, r, nil)
	client := NewClient(socket, id, username, hub)

	hub.Register(client)
}
