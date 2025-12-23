package realtime

import "github.com/gorilla/websocket"

type Client struct {
	// connection
	Conn *websocket.Conn
	// channel for server ---> client message
	Send chan []byte
	// rooms    every image id on screen is room for this client
	Rooms map[string]bool
}

// join and leave rooms
func (c *Client) ReadPump(hub *Hub) {

	defer func() {
		hub.Unregister <- c
		c.Conn.Close()
	}()

	for {
		var msg map[string]string
		err := c.Conn.ReadJSON(&msg)
		if err != nil {
			break
		}

		if msg["action"] == "JOIN_ROOM" {

			room := msg["roomId"]
			c.Rooms[room] = true

			if hub.RoomsHub[room] == nil {
				hub.RoomsHub[room] = make(map[*Client]bool)
			}

			hub.RoomsHub[room][c] = true
		}
	}

}

// write pump
func (c *Client) WritePump() {
	defer c.Conn.Close()
	for msg := range c.Send {
		c.Conn.WriteMessage(websocket.TextMessage, msg)
	}
}
