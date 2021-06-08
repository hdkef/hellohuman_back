package main

import (
	"database/sql"
	"fmt"
	"hellohuman/controller"
	"hellohuman/static"
	"hellohuman/utils"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

//Load the environment variable
func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {

	login := controller.NewLoginHandler(&sql.DB{})
	room := controller.NewRoomHandler(&sql.DB{})

	router := mux.NewRouter()
	router.HandleFunc(static.LoginRoute, utils.Cors(login.Login()))
	router.HandleFunc(static.EstablishWSRoute, room.EstablishWS())

	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))

	fmt.Println("about to listen and serve")

	http.ListenAndServe(addr, router)
}
