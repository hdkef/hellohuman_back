package models

import "github.com/gorilla/websocket"

type User struct {
	ID     string
	RoomID string
	Conn   *websocket.Conn
	Name   string
	Gender string
	SDP    interface{}
	ICE    []interface{}
}
