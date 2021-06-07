package controller

import (
	"hellohuman/models"
	"hellohuman/static"
)

//answerFromClient handle new sdp answer from client
func answerFromClient(payload models.WSPayload) {
	go insertUserSDP(&payload.User, &payload.SDP)
	go sendStoredICE(&payload.User)
}

func sendStoredICE(user *models.User) {
	var ices []interface{}
	for _, v := range rooms[user.RoomID] {
		if v.Conn != user.Conn {
			ices = v.ICE
			break
		}
	}

	if len(ices) == 0 {
		return
	}

	for _, v := range ices {
		resp := models.ICEResponse{
			Type: static.ICEFromServer,
			ICE:  v,
		}
		user.Conn.WriteJSON(resp)
	}
}
