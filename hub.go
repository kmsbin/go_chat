package main

type hub struct {
	Registered   map[string]client
	Unregistered map[string]client
}

func newHub() *hub {
	return &hub{
		Registered:   make(map[string]client, 10),
		Unregistered: make(map[string]client, 10),
	}
}

func (hub *hub) Register(newClient client) {
	hub.Registered[newClient.Id] = newClient
}
func (hub *hub) Unregister(client client) {
	delete(hub.Registered, client.Id)
	hub.Unregistered[client.Id] = client
}
