package handlers

import (
	"fmt"
	"runner/types"

	"github.com/gorilla/websocket"
)

func FileTree(msg types.Message, conn *websocket.Conn) {
	fmt.Println(msg)
}
