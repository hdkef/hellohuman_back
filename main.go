package main

import (
	"database/sql"
	"fmt"
	"hellohuman/controller"
	"hellohuman/static"
	"hellohuman/utils"
	"net/http"
	"os"
	"path/filepath"

	"github.com/NYTimes/gziphandler"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type spaHandler struct {
	staticPath string
	indexPath  string
}

//serveHTTP will serve public html, javascript, css, and images
func (h spaHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	path, err := filepath.Abs(req.URL.Path)
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}

	path = filepath.Join(h.staticPath, req.URL.Path) // literally pointing to /dist/angular/{{the url}}

	_, err = os.Stat(path)
	if os.IsNotExist(err) { //if the requested path to file is not available, give index.html
		http.ServeFile(res, req, filepath.Join(h.staticPath, h.indexPath))
		return
	} else if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	http.FileServer(http.Dir(h.staticPath)).ServeHTTP(res, req) //if the requested path available, serve that file relative to staticPath
}

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

	spa := spaHandler{staticPath: os.Getenv("STATICPATH"), indexPath: "index.html"}
	router.PathPrefix("/").Handler(spa)

	addr := fmt.Sprintf(":%s", os.Getenv("PORT"))

	fmt.Println("about to listen and serve")

	http.ListenAndServe(addr, gziphandler.GzipHandler(router))
}
