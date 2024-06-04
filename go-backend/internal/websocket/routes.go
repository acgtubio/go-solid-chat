package websocket

import (
	"net/http"

	"github.com/acgtubio/go-solid-chat/internal/chat"
)

func NewChatServer() http.Handler {
	chatRoom := chat.CreateRooms()
	mux := http.NewServeMux()

	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		ChatHandler(chatRoom, w, r)
	})
	mux.HandleFunc("/test", TestHandler)

	return mux
}

func createChatRooms() *chat.Rooms {
	rooms := chat.CreateRooms()
	rooms.Start()

	defaultRoom := chat.Room{
		ID:         "0",
		Name:       "Default Channel",
		Register:   make(chan *chat.Client),
		Unregister: make(chan string),
		Broadcast:  make(chan chat.Message),
		Members:    []chat.Client{},
	}

	rooms.RegisterRoom <- &defaultRoom

	return rooms
}
