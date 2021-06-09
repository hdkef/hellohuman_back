package controller

import (
	"hellohuman/models"
	"hellohuman/static"
	"time"
)

//chatFromClient is to handle chat request from client
func chatFromClient(payload models.WSPayload) {
	ws, exist := onlineMap[payload.Peer.ID]
	if exist != true {
		//Handle peer is offline
		return
	}
	resp := models.ChatResponse{
		Type:   static.ChatFromPeer,
		Name:   payload.User.Name,
		Gender: payload.User.Name,
		Text:   payload.Text,
		Date:   time.Now().String(),
	}
	ws.WriteJSON(resp)
	go chatFromMe(&payload)
}

//chatFromMe the chat itself will be sent back to client to know whether the chat is succeed or not
func chatFromMe(payload *models.WSPayload) {
	resp := models.ChatResponse{
		Type:   static.ChatFromPeer,
		Name:   payload.User.Name,
		Gender: payload.User.Name,
		Text:   payload.Text,
		Date:   time.Now().String(),
	}
	payload.User.Conn.WriteJSON(resp)
}
