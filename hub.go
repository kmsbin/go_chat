package main

type Hub struct {
	Registered   map[string]client
	Unregistered map[string]client
	History      map[string]string
}

func newHub() Hub {
	return Hub{
		Registered:   make(map[string]client, 10),
		Unregistered: make(map[string]client, 10),
		History:      make(map[string]string, 10),
	}
}

func (Hub *Hub) Register(syncClient client) string {
	go syncClient.ReadMessagePool()
	go syncClient.WriteMessagePool()
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
	}
}
