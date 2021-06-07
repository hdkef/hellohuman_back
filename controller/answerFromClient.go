package controller

import "hellohuman/models"

//answerFromClient handle new sdp answer from client
func answerFromClient(payload models.WSPayload) {
	go insertUserSDP(&payload.User, &payload.SDP)
}
