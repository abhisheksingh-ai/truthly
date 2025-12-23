package controller

import (
	"net/http"
	"truthly/internals/realtime"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ServeWS(hub *realtime.Hub, w http.ResponseWriter, r *http.Request) {
	// upgrade HTTP to WebSocket
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "WebSocket upgrade failed", http.StatusBadRequest)
		return
	}

	client := &realtime.Client{
		Conn:  conn,
		Send:  make(chan []byte, 256),
		Rooms: make(map[string]bool),
	}

	hub.Register <- client

	go client.WritePump()
	go client.ReadPump(hub)
}
