package models

//WSPayload is how websocket payload formatted
type WSPayload struct {
	Type string
	User User
	SDP  interface{}
	ICE  interface{}
}

//LoginPayload is how login request's body data formatted
type LoginPayload struct {
	User User
}
