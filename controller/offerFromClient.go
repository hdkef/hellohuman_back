package controller

import "hellohuman/models"

//offerFromClient handle new SDP offer from client
func offerFromClient(payload models.WSPayload) {
	go insertUserSDP(&payload.User, &payload.SDP)
}

//insertUserSDP will insert client's sdp to corresponding user sdp field for roomID in rooms map
func insertUserSDP(user *models.User, SDP *interface{}) {
	for _, v := range rooms[user.RoomID] {
		if (*v).ID == user.ID {
			v.SDP = SDP
			return
		}
	}
}
