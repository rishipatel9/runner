package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port := getPort()

	server := &http.Server{
		Addr: fmt.Sprintf(":%s", port),
	}

	http.HandleFunc("/", handleWebSocket)

	fmt.Println("Server running on port 4000")
	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}
	return port
}
