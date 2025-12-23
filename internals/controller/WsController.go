package controller

import (
	"net/http"
	"truthly/internals/realtime"
	"truthly/internals/util/auth"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ServeWS(hub *realtime.Hub, w http.ResponseWriter, r *http.Request, authToken *auth.AuthToken) {
	// upgrade HTTP to WebSocket
	//  1. extract token
	token := r.URL.Query().Get("token")
	if token == "" {
		http.Error(w, "token required", http.StatusUnauthorized)
		return
	}

	//  2. verify token
	claims, err := authToken.VerifyJwtToken(token, r.Context())
	if err != nil {
		http.Error(w, "invalid token", http.StatusUnauthorized)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		http.Error(w, "WebSocket upgrade failed", http.StatusBadRequest)
		return
	}

	client := &realtime.Client{
		Conn:   conn,
		Send:   make(chan []byte, 256),
		Rooms:  make(map[string]bool),
		UserId: claims.UserId,
	}

	hub.Register <- client

	go client.WritePump()
	go client.ReadPump(hub)
}
