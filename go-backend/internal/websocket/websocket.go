package websocket

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/acgtubio/go-solid-chat/internal/chat"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func ChatHandler(chatRoom *chat.Rooms, w http.ResponseWriter, r *http.Request) {
	ws, err := upgrade(w, r)

	if err != nil {
		fmt.Fprintf(w, "%+V\n", err)
	}

	// reader(ws)
	client := chat.Client{
		ID:   uuid.NewString(),
		Conn: ws,
	}
}

func upgrade(w http.ResponseWriter, r *http.Request) (*websocket.Conn, error) {

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return ws, err
	}

	return ws, nil
}

func reader(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()

		if err != nil {
			log.Println(err)
			return
		}

		var m chat.MessageBody
		err = json.Unmarshal(p, &m)
		fmt.Println(m.Content)

		fmt.Println(messageType)

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}
