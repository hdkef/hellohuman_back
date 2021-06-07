package utils

import (
	"encoding/json"
	"net/http"
)

//ResOK send ok response with MSG payload
func ResOK(res *http.ResponseWriter, msg string) {
	obj := struct {
		MSG string
	}{
		MSG: msg,
	}
	(*res).WriteHeader(http.StatusOK)
	_ = json.NewEncoder(*res).Encode(&obj)
}
