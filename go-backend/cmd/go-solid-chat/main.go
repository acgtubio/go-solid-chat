package main

import (
	"fmt"
	"net/http"

	"github.com/acgtubio/go-solid-chat/internal/websocket"
)

func main() {
	router := websocket.NewChatServer()
	fmt.Printf("Server running and listening on localhost:8080\n")
	err := http.ListenAndServe(":8080", router)

	if err != nil {
		panic(err)
	}
}
