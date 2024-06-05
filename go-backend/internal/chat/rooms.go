package chat

type Rooms struct {
	RegisterRoom   chan *Room
	RoomCollection []*Room
}

func CreateRooms() *Rooms {
	return &Rooms{
		RegisterRoom:   make(chan *Room),
		RoomCollection: []*Room{},
	}
}

func (r *Rooms) Start() {
	for {
		select {
		case room := <-r.RegisterRoom:
			r.RoomCollection = append(r.RoomCollection, room)
		}
	}
}
