package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	http.HandleFunc("/ws", wsHandler)

	http.ListenAndServe(":8080", nil)
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "Failed to upgrade to websocket", http.StatusBadRequest)
		fmt.Println("Upgrader failed")
		return
	}
	defer conn.Close()

	fmt.Println("New WebSocket: 连接成功")

	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("ReadMessage failed")
			return
		}
		fmt.Println("Received messageType:", messageType)
		err = conn.WriteMessage(messageType, message)
		message = append(message, "已收到"...)
		if err != nil {
			fmt.Println("WriteMessage failed")
			return
		}

	}
}
