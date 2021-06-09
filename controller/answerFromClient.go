package controller

import (
	"hellohuman/models"
	"hellohuman/static"
)

//answerFromClient handle new sdp answer from client
func answerFromClient(payload models.WSPayload) {
	go sendAnswer(&payload.User, &payload.SDP)
	go insertUserSDP(&payload.User, &payload.SDP)
	go sendStoredICE(&payload.User)
}

//sendAnswer will give answer to caller client and peerinfo
func sendAnswer(user *models.User, SDP *interface{}) {
	resp := models.AnswerResponse{
		Type: static.AnswerFromServer,
		SDP:  SDP,
		Peer: *user,
	}
	for _, v := range rooms[user.RoomID] {
		if v.Conn != user.Conn {
			v.Conn.WriteJSON(resp)
			break
		}
	}
}

//sendStoredICE will send the stored ICE to user, basically stored ICE from caller client
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
