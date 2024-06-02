package chat

type Message struct {
	RoomID string `json:"roomId"`
	Author string `json:"author"`
	Type   string `json:"type"`
	Body   string `json:"body"`
}
