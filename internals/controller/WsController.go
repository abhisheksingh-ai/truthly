package controller

import (
	"log/slog"
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

func ServeWS(
	hub *realtime.Hub,
	w http.ResponseWriter,
	r *http.Request,
	authToken *auth.AuthToken,
	logger *slog.Logger,
) {
	// Upgrade FIRST
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	logger.Info("Connection upgraded to ws")

	// Extract token
	token := r.URL.Query().Get("token")
	if token == "" {
		conn.WriteMessage(websocket.CloseMessage,
			websocket.FormatCloseMessage(
				websocket.ClosePolicyViolation,
				"token required",
			),
		)
		logger.Error("ws conn closed, token missing")
		conn.Close()
		return
	}

	//Verify token
	claims, err := authToken.VerifyJwtToken(token, r.Context())
	if err != nil {
		conn.WriteMessage(websocket.CloseMessage,
			websocket.FormatCloseMessage(
				websocket.ClosePolicyViolation,
				"invalid token",
			),
		)
		conn.Close()
		return
	}

	//Create client
	client := &realtime.Client{
		Conn:   conn,
		Send:   make(chan []byte, 256),
		Rooms:  make(map[string]bool),
		UserId: claims.UserId,
	}

	//Register client
	hub.Register <- client

	//Start pumps
	go client.WritePump()
	go client.ReadPump(hub)
}
