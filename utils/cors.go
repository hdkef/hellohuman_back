package utils

import "net/http"

//handleCors is handle cross origin resource
func handleCors(res *http.ResponseWriter) {
	(*res).Header().Set("Access-Control-Allow-Origin", "*")
	(*res).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*res).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Accept-Encoding, X-CSRF-Token, Bearer")
}

//Cors is the middleware for cors AND preflight, next will be executed after cors
func Cors(next http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		handleCors(&res)
		if option := handlePreflight(&res, req); option == true {
			return
		}
		next.ServeHTTP(res, req)
	}
}
