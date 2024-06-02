package websocket

import (
	"net/http"
)

func NewChatServer() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/ws", ChatHandler)
	mux.HandleFunc("/test", TestHandler)

	return mux
}
