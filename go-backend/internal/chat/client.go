package chat

import (
	"encoding/json"
	"log"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

type Client struct {
	ID    string
	Conn  *websocket.Conn
	Rooms []*Room
}

func NewClient(ws *websocket.Conn) *Client {
	return &Client{
		ID:    uuid.NewString(),
		Conn:  ws,
		Rooms: []*Room{},
	}
}

// Reads messages from websocket connection. Message received from the client
// will be parsed into the MessageBody struct, then send the message to all the clients in the room.
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
		message := Message{
			Type: messageType,
			Body: messageBody,
		}

		roomID := messageBody.RoomID
		for _, r := range c.Rooms {
			if r.ID == roomID {
				r.Broadcast <- message
				break
			}
		}
	}
}

// Susbscribes the client to the room, and adds the room to the list of rooms the client is connected to.
func (c *Client) JoinRoom(r *Room) {
	r.Register <- c
	c.Rooms = append(c.Rooms, r)
}
