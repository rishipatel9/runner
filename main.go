package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.Handle("/", http.HandlerFunc(createHandler)).Methods("GET")

	fmt.Println("Server running on port 4000")
	log.Fatal(http.ListenAndServe(":4000", r))
}

func createHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Create handler called"))
}
