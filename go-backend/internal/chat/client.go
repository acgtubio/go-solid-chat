package chat

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	ID    string
	Conn  *websocket.Conn
	Rooms []*Room
}

func (c *Client) Read() {
	defer func() {
		c.Conn.Close()
	}()

	for {
		var messageBody MessageBody
		messageType, body, err := c.Conn.ReadMessage()

		if err != nil {
			log.Println(err)
			// TODO: Create functionality to return this to client to indicate message has not been sent.
			return
		}

		err = json.Unmarshal(body, &messageBody)
		if err != nil {
			log.Println(err)
			return
		}
		room := messageBody.RoomID

		message := Message{
			Type: messageType,
			Body: messageBody,
		}

		for _, r := range c.Rooms {
			if r.ID == room {
				r.Broadcast <- message
				break
			}
		}
		fmt.Println(message)
	}
}

func (c *Client) JoinRoom(r *Room) {
	c.Rooms = append(c.Rooms, r)
}
