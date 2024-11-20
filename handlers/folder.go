package handlers

import (
	"fmt"
	"runner/types"

	"github.com/gorilla/websocket"
)

func Folder(msg types.Message, conn *websocket.Conn) {
	fmt.Println(msg)
}
