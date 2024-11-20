package handlers

import (
	"net/http"
)

func CreateHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Create handler invoked"))
}
