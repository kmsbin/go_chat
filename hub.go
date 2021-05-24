package main

type hub struct {
	Registered   map[string]client
	Unregistered map[string]client
}

func (hub *hub) Register(newClient client) {
	hub.Registered[newClient.Id] = newClient
}
func (hub *hub) Unregister(client client) {
	delete(hub.Registered, client.Id)
	hub.Unregistered[client.Id] = client
}
