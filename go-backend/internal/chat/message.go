package chat

type Message struct {
	Type int         `json:"type"`
	Body MessageBody `json:"body"`
}

type MessageBody struct {
	Type    string `json:"messageType"`
	Author  string `json:"author"`
	RoomID  string `json:"roomId"`
	Content string `json:"content"`
}
