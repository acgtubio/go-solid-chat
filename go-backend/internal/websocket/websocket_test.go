package websocket

import (
	"testing"
)

func TestCreateRooms(t *testing.T) {
	rooms := CreateChatRooms()

	length := len(rooms.RoomCollection)
	wantLength := 1

	if length != wantLength {
		t.Errorf("got %q, wanted %q", length, wantLength)
	}
}
