package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"runner/handlers"
	"runner/types"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("Failed to upgrade connection:", err)
		return
	}
	defer conn.Close()

	fmt.Println("New Client Connection")

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Print("Error reading msg", err)
		}
		var message types.Message
		if err := json.Unmarshal(msg, &message); err != nil {
			log.Println("Invalid message format:", err)
			continue
		}
		fmt.Println(message)

		switch message.Type {
		case "file":
			handlers.File(message, conn)
		case "folder":
			handlers.Folder(message, conn)
		case "fetch":
			handlers.Fetch(message, conn)
		case "filetree":
			handlers.FileTree(message, conn)
		case "terminal":
			handlers.Terminal(message, conn)
		default:
			conn.WriteJSON(map[string]string{"message": "Unknown message type"})
		}

	}

}
