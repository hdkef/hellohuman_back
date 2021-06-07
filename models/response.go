package models

//LoginResponse format login response json
type LoginResponse struct {
	User User
}

//RoomResponse format createRoom response or joinRoom response to client
type RoomResponse struct {
	Type   string
	RoomID string
}
