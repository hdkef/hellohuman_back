package utils

import "net/http"

//handlePreflight handle pre flight options
func handlePreflight(res *http.ResponseWriter, req *http.Request) bool {
	if req.Method == http.MethodOptions {
		(*res).WriteHeader(http.StatusOK)
		return true
	} else {
		return false
	}
}

//Preflight is middleware to handle pre flight options
func Preflight(next http.HandlerFunc) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		if option := handlePreflight(&res, req); option == true {
			return
		}
		next.ServeHTTP(res, req)
	}
}
