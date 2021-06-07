package controller

import (
	"database/sql"
	"encoding/json"
	"hellohuman/models"
	"hellohuman/utils"
	"net/http"

	"github.com/google/uuid"
)

//LoginHandle is a struct contains all data needed for login handler
type LoginHandle struct {
	DB *sql.DB
}

//NewLoginHandler will return new pointer for loginhandle
func NewLoginHandler(db *sql.DB) *LoginHandle {
	return &LoginHandle{DB: db}
}

//Login will handle login request
func (x *LoginHandle) Login() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {

		var loginpayload models.LoginPayload

		err := json.NewDecoder(req.Body).Decode(&loginpayload)
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}

		loginresponse := models.LoginResponse{
			ID:     uuid.New().String(),
			Name:   loginpayload.Name,
			Gender: loginpayload.Gender,
		}

		err = json.NewEncoder(res).Encode(&loginresponse)
		if err != nil {
			utils.ResErr(&res, http.StatusInternalServerError, err)
			return
		}
	}
}
