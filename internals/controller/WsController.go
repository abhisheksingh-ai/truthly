package controller

import (
	"fmt"
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
) {
	// 1️⃣ Upgrade FIRST
	fmt.Print("Hey")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}

	// 2️⃣ Extract token
	token := r.URL.Query().Get("token")
	if token == "" {
		conn.WriteMessage(websocket.CloseMessage,
			websocket.FormatCloseMessage(
				websocket.ClosePolicyViolation,
				"token required",
			),
		)
		conn.Close()
		return
	}

	// 3️⃣ Verify token
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

	// 4️⃣ Create client
	client := &realtime.Client{
		Conn:   conn,
		Send:   make(chan []byte, 256),
		Rooms:  make(map[string]bool),
		UserId: claims.UserId,
	}

	// 5️⃣ Register client
	hub.Register <- client

	// 6️⃣ Start pumps
	go client.WritePump()
	go client.ReadPump(hub)
}
