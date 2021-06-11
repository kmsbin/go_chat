package main

type Hub struct {
	Registered   map[string]Client `json:"registered-user,omitempty"`
	Unregistered map[string]Client `json:"unregistered-user,omitempty"`
	History      map[string]string `json:"-"`
}

func newHub() Hub {
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
func (Hub *Hub) Unregister(clientChan chan Client) {
	newClient := <-clientChan

	delete(Hub.Registered, newClient.Id)
	Hub.Unregistered[newClient.Id] = newClient
}
func (Hub *Hub) Run() {
	// for {
	// for _, newClient := range Hub.Registered {
	// 	msgType, msg, err := newClient.Socket.ReadMessage()

	// 	if err != nil {
	// 		log.Println(err)
	// 		return
	// 	}
	// 	var message Message
	// 	_ = json.Unmarshal(msg, &message)
	// 	log.Println(message.Message)
	// 	Hub.Registered[message.IdReciever].Socket.WriteMessage(msgType, []byte(message.Message))
	// 	// log.Println("Mensagem recebida: ", msg)
	// }
	// }
}
