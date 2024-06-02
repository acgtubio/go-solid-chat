package chat

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID    string
	Conn  *websocket.Conn
	Rooms []Room
}

func (c *Client) Read() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		_, body, err := c.Conn.ReadMessage()

		if err != nil {
			log.Println(err)

			// TODO: Create functionality to return this to client to indicate message has not been sent.
			return
		}

		message := Message{
			Type:   "Test",
			Author: c.ID,
			RoomID: "Default",
			Body:   string(body),
		}

		fmt.Println(message)
	}
}
