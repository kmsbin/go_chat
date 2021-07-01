package entities

type Message struct {
	IdReciever string `json:"id-reciever"`
	Message    string `json:"message"`
	IdSender   string `json:"id-sender,omitempty"`
	Username   string `json:"username,omitempty"`
}
