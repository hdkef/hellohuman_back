package controller

import (
	"context"
	"database/sql"
	"hellohuman/models"
	"hellohuman/static"
	"hellohuman/utils"
	"net/http"

	"github.com/gorilla/websocket"
)

//RoomHandle is a struct contains all data needed for room
type RoomHandle struct {
	DB *sql.DB
}

//NewRoomHandler return new pointer of Room Handle
func NewRoomHandler(db *sql.DB) *RoomHandle {
	return &RoomHandle{DB: db}
}

//store the user info to map, to access this need RoomID
var rooms map[string][]*models.User = make(map[string][]*models.User)

//upgrader to upgrade http protocol to websocket
var upgrader *websocket.Upgrader = &websocket.Upgrader{
	CheckOrigin: func(req *http.Request) bool {
		return true
	},
}

//EstablishWS is where the websocket handshake happens between client and server
func (x *RoomHandle) EstablishWS() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

		ws, err := upgrader.Upgrade(res, req, res.Header())
		if err != nil {
			utils.ResErr(&res, http.StatusBadRequest, err)
			return
		}

		ctx, cancel := context.WithCancel(context.Background())
		payloadChan := make(chan models.WSPayload)
		go payloadReader(cancel, ws, x.DB, payloadChan)
		go payloadRouter(ctx, payloadChan)
	}
}

//payloadReader will create one goroutine that read all incoming payload, append a ws pointer, and give it to payloadChan
func payloadReader(cancel context.CancelFunc, ws *websocket.Conn, DB *sql.DB, payloadChan chan models.WSPayload) {

	var payload models.WSPayload
	payload.User.Conn = ws

	defer cancel()

	for {
		err := websocket.ReadJSON(ws, &payload)
		if err != nil {
			close(payloadChan)
			break //if read is error which means connection is disconnected, then break from loop and cancel()
		}
		payloadChan <- payload
	}
}

//payloadRouter will receive payload from channel and route it according to payload type
func payloadRouter(ctx context.Context, payloadChan chan models.WSPayload) {
	for {
		select {
		case <-ctx.Done():
			return
		case payload := <-payloadChan:
			switch payload.Type {
			case static.InitFromClient:
				go initFromClient(payload)
			case static.OfferFromClient:
				go offerFromClient(payload)
			case static.AnswerFromClient:
				go answerFromClient(payload)
			case static.ICEFromClient:
				go ICEFromClient(payload)
			}
		}
	}
}
