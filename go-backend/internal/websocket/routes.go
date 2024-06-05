package websocket

import (
	"net/http"

	"github.com/acgtubio/go-solid-chat/internal/chat"
)

func NewChatServer() http.Handler {
	chatRoom := CreateChatRooms()
	mux := http.NewServeMux()

	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ChatHandler(chatRoom, w, r)
	})
	mux.HandleFunc("/test", TestHandler)

	return mux
}

func CreateChatRooms() *chat.Rooms {
	rooms := chat.CreateRooms()
	go rooms.Start()

	defaultRoom := chat.Room{
		ID:         "0",
		Name:       "Default Channel",
		Register:   make(chan *chat.Client),
		Unregister: make(chan string),
		Broadcast:  make(chan chat.Message),
		Members:    []*chat.Client{},
	}
	go defaultRoom.Start()

	rooms.RegisterRoom <- &defaultRoom

	return rooms
}
