package models

import "github.com/gorilla/websocket"

//WSPayload is how websocket payload formatted
type WSPayload struct {
	Conn *websocket.Conn
	Type string
}

//LoginPayload is how login request's body data formatted
type LoginPayload struct {
	Name   string
	Gender string
}
