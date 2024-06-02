package chat

import "fmt"

type Room struct {
	ID         string
	Name       string
	Register   chan *Client
	Unregister chan string
	Broadcast  chan Message
	Members    []Client
}

func (r *Room) Start() {
	for {
		select {
		case client := <-r.Register:
			r.Members = append(r.Members, *client)
		case message := <-r.Broadcast:
			for _, client := range r.Members {
				if err := client.Conn.WriteJSON(message); err != nil {
					fmt.Println(err)
					return
				}
			}
		}
	}
}

func NewRoom(id, name string) *Room {
	return &Room{
		ID:         id,
		Name:       name,
		Register:   make(chan *Client),
		Unregister: make(chan string),
		Broadcast:  make(chan Message),
		Members:    []Client{},
	}
}
