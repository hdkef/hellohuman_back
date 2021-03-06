package models

//LoginResponse format login response json
type LoginResponse struct {
	User User
}

//RoomResponse format createRoom response or joinRoom response to client
type RoomResponse struct {
	Type   string
	RoomID string
	SDP    interface{}
	Peer   User
}

//ICEResponse format ICEFromServer response
type ICEResponse struct {
	Type string
	ICE  interface{}
}

//AnswerResponse format answerfromserver
type AnswerResponse struct {
	Type string
	SDP  interface{}
	Peer User
}

//DisconnectResponse how data when user disconnected and want to tell peer formatted
type DisconnectResponse struct {
	Type string
}

//ChatResponse how chat data is formatted
type ChatResponse struct {
	Name   string
	Gender string
	Type   string
	Text   string
	Date   string
}
