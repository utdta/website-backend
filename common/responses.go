package common

import (
	"net/http"
)

func WriteJson(s int, j []byte, w http.ResponseWriter) {
	w.Header().Set(`Content-Type`, `application/json`)
	w.WriteHeader(s)
	w.Write(j)
}
