package handlers

import (
	"fmt"
	"runner/types"

	"github.com/gorilla/websocket"
)

func File(message types.Message, conn *websocket.Conn) {
	fmt.Println(message)

}
