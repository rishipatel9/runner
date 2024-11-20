package handlers

import (
	"encoding/json"
	"log"
	"runner/terminal"
	"runner/types"

	"github.com/gorilla/websocket"
)

var termManager = terminal.NewTerminal()

func Terminal(msg types.Message, conn *websocket.Conn) {
	var actionData map[string]interface{}
	if err := json.Unmarshal([]byte(msg.Payload.Action), &actionData); err != nil {
		log.Println("Invalid payload format:", err)
		conn.WriteMessage(websocket.TextMessage, []byte("Invalid payload format"))
		return
	}

	switch actionData["action"] {
	case "init":
		id := actionData["id"].(string)
		replID := actionData["replId"].(string)
		onData := func(data string) {
			conn.WriteMessage(websocket.TextMessage, []byte(data))
		}
		_, err := termManager.Init(id, replID, onData)
		if err != nil {
			log.Println("Failed to initialize terminal:", err)
			conn.WriteMessage(websocket.TextMessage, []byte("Error initializing terminal"))
		} else {
			conn.WriteMessage(websocket.TextMessage, []byte("Terminal initialized successfully"))
		}

	case "write":
		id := actionData["id"].(string)
		data := actionData["data"].(string)
		termManager.Write(id, data)

	case "resize":

		id := actionData["id"].(string)
		cols := int(actionData["cols"].(float64))
		rows := int(actionData["rows"].(float64))
		termManager.Resize(id, cols, rows)

	case "close":

		id := actionData["id"].(string)
		termManager.Close(id)

	default:
		log.Println("Unknown action:", actionData["action"])
		conn.WriteMessage(websocket.TextMessage, []byte("Unknown action"))
	}
}
