package controller

import (
	"hellohuman/models"
	"hellohuman/static"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

//initFromClient handle when the client send payload 'initFromClient' which is the first thing after conn is established
func initFromClient(payload models.WSPayload) {

	//Here code to response initFromClient
	empty := isRoomsEmpty()
	if empty == true {
		go createRoom(&payload.User)
		return
	}
	roomID := isRoomAvail()
	if roomID == "" {
		go createRoom(&payload.User)
		return
	}

	go joinRoom(&payload.User, &roomID)
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
func createRoom(user *models.User) {
	roomID := "room" + uuid.New().String()
	rooms[roomID] = []*models.User{user}
	onlineMap[user.ID] = user.Conn
	resp := models.RoomResponse{
		Type:   static.CreatedRoomFromServer,
		RoomID: roomID,
	}
	go pingPonger(user, &roomID)
	user.Conn.WriteJSON(resp)
}

//joinRoom will append new ws to roomID and give offer and peer info
func joinRoom(user *models.User, roomID *string) {
	rooms[*roomID] = append(rooms[*roomID], user)
	onlineMap[user.ID] = user.Conn
	var peer *models.User
	for _, v := range rooms[*roomID] {
		if v.Conn != user.Conn {
			peer = v
			break
		}
	}
	resp := models.RoomResponse{
		Type:   static.JoinedRoomFromServer,
		RoomID: *roomID,
		SDP:    peer.SDP,
		Peer:   *peer,
	}
	go pingPonger(user, roomID)
	user.Conn.WriteJSON(resp)
}

const (
	pongWait   = 5 * time.Second
	pingPeriod = (pongWait * 9) / 10
)

//pingPonger will create one goroutine that ping the client, and when the connection is lost. It deletes all related to client's online trace (ID)
//roomID comes directly from roomID not from user.RoomID because InitFromClient doesn't bring room
//User.RoomID only available in offerFromClient and answerFromClient
func pingPonger(user *models.User, roomID *string) {
	user.Conn.SetPongHandler(func(appData string) error {
		user.Conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	timer := time.NewTicker(pingPeriod)

	defer func(user *models.User) {
		timer.Stop()

		if _, exist := onlineMap[user.ID]; exist == true {
			delete(onlineMap, user.ID)
		}

		if len(rooms[*roomID]) == 1 {
			delete(rooms, *roomID) //if the room is only 1 person, delete the map
			return
		}

		deletedUsers := []*models.User{}
		for _, v := range rooms[*roomID] { //if the room is 2 person, only delete user that disconnected
			if v.Conn != user.Conn {
				deletedUsers = append(deletedUsers, v)
				go tellPeerMeDisconnect(v.Conn)
				break
			}
		}
		rooms[*roomID] = deletedUsers
	}(user)

	for {
		select {
		case <-timer.C:
			if err := user.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

//tellPeerMeDisconnect will send to peer client that user disconnected
func tellPeerMeDisconnect(ws *websocket.Conn) {
	resp := models.DisconnectResponse{
		Type: static.PeerDisconnected,
	}
	ws.WriteJSON(resp)
}
