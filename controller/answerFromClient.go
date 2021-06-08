package controller

import (
	"hellohuman/models"
	"hellohuman/static"
)

//answerFromClient handle new sdp answer from client
func answerFromClient(payload models.WSPayload) {
	go getOfferSendAnswer(&payload.User)
	go insertUserSDP(&payload.User, &payload.SDP)
	go sendStoredICE(&payload.User)
}

//getAnswerSendOffer will try to get offer for user and send answer to caller client
func getOfferSendAnswer(user *models.User) {
	resp := models.OfferAnswerResponse{
		Type: static.AnswerFromServer,
		SDP:  user.SDP,
	}
	var otherSDP interface{}
	for _, v := range rooms[user.RoomID] {
		if v.Conn != user.Conn {
			otherSDP = v.SDP
			v.Conn.WriteJSON(resp)
			break
		}
	}
	resp.Type = static.OfferFromServer
	resp.SDP = otherSDP
	user.Conn.WriteJSON(resp)
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
