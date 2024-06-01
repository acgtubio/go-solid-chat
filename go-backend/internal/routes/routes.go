package routes

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func NewChatServer() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/ws", wsUpgrade)
	mux.HandleFunc("/test", testHandler)

	return mux
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "This is the test route")
}

func wsUpgrade(w http.ResponseWriter, r *http.Request) {

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)

		// TODO: Determine if return is needed. I put this here because it is intuitive to return if there is an error.
		return
	}

	reader(ws)
}

func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()

		if err != nil {
			log.Println(err)
			return
		}

		fmt.Println(string(p))

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}
