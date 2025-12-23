package realtime

import "github.com/gorilla/websocket"

type Client struct {
	// connection
	conn *websocket.Conn
	// channel for server ---> client message
	send chan []byte
	// rooms    every image id on screen is room for this client
	rooms map[string]bool
}
