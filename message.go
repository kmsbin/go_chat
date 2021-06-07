package main

type Message struct {
	IdSender   string `json:"id-sender"`
	IdReciever string `json:"id-reciever"`
	Username   string `json:"username"`
	Message    string `json:"message"`
	MsgType    int    `json:"msgtype,omitempty"`
}
