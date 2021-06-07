package controller

import (
	"database/sql"
	"net/http"
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
		res.Write([]byte("Login"))
	}
}
