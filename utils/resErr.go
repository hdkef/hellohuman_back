package utils

import (
	"fmt"
	"net/http"
)

//Respond with error
func ResErr(res *http.ResponseWriter, code int, err error) {
	payload := fmt.Sprintf("%s", err.Error())
	(*res).WriteHeader(code)
	(*res).Write([]byte(payload))
}
