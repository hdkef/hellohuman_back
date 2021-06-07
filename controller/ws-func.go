package controller

import (
	"hellohuman/models"
	"hellohuman/static"

	"github.com/google/uuid"
)

//initFromClient handle when the client send payload 'initFromClient' which is the first thing after conn is established
func initFromClient(payload models.WSPayload) {

	//Here code to response initFromClient
	empty := isRoomsEmpty()
	if empty == true {
		go createRoom(payload.User)
		return
	}
	roomID := isRoomAvail()
	if roomID == "" {
		go createRoom(payload.User)
		return
	}

	go joinRoom(payload.User, roomID)
}

//isRoomsEmpty return true if the rooms map empty
func isRoomsEmpty() bool {
	return len(rooms) == 0
}

//isRoomAvail return roomID if the room is available or return "" if full
func isRoomAvail() string {
	for s, v := range rooms {
		if len(v) == 1 {
			return s
		}
	}
	return ""
}

//createRoom will add new room and add user's info and respond roomID
func createRoom(user models.User) {
	roomID := "room" + uuid.New().String()
	rooms[roomID] = []*models.User{&user}
	resp := models.RoomResponse{
		Type:   static.CreatedRoomFromServer,
		RoomID: roomID,
	}
	go pingPonger(user, roomID)
	user.Conn.WriteJSON(resp)
}

//joinRoom will append new ws to roomID and respond roomID
func joinRoom(user models.User, roomID string) {
	rooms[roomID] = append(rooms[roomID], &user)
	resp := models.RoomResponse{
		Type:   static.JoinedRoomFromServer,
		RoomID: roomID,
	}
	go pingPonger(user, roomID)
	user.Conn.WriteJSON(resp)
}

//pingPonger will create one goroutine that ping the client, and when the connection is lost. It deletes all related to client's online trace (ID)
func pingPonger(user models.User, roomID string) {

}
