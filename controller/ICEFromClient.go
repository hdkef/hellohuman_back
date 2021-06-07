package controller

import (
	"hellohuman/models"
	"hellohuman/static"
)

//ICEFromClient handle new ice from client
func ICEFromClient(payload models.WSPayload) {
	go addUserICE(&payload.User, &payload.ICE)
	go sendUserICE(&payload.User, &payload.ICE)
}

//addUserICE will add user new ice to corresponding user ice array at roomID in rooms map
func addUserICE(user *models.User, ICE *interface{}) {
	for _, v := range rooms[user.RoomID] {
		if v.ID == user.ID {
			v.ICE = append(v.ICE, ICE)
			break
		}
	}
}

//sendUserICE will send ICE to other user in same roomID
func sendUserICE(user *models.User, ICE *interface{}) {
	resp := models.ICEResponse{
		Type: static.ICEFromServer,
		ICE:  ICE,
	}
	for _, v := range rooms[user.RoomID] {
		if v.Conn != user.Conn {
			v.Conn.WriteJSON(resp)
			break
		}
	}
}
