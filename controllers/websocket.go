package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

var clients = make(map[*websocket.Conn]bool)

func WebSocketHandler(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer conn.Close()
	clients[conn] = true
	for {
		_, _, err := conn.ReadMessage()
		if err != nil {
			delete(clients, conn)
			return
		}
	}
}

func Broadcast(message string) {
	for client := range clients {
		client.WriteMessage(websocket.TextMessage, []byte(message))
	}
}
