package realtime

import "encoding/json"

type Hub struct {
	// clients
	clients map[*Client]bool
	// rooms
	rooms map[string]map[*Client]bool

	// register
	register chan *Client

	// unregister
	unregister chan *Client

	//broadcast channer
	broadcast chan Event
}

func (h *Hub) Run() {
	for {
		select {

		case client := <-h.register:
			h.clients[client] = true

		case client := <-h.unregister:
			// clients se hata do
			delete(h.clients, client)
			for room := range client.rooms {
				// Jis bhi image id se ye client connected h vha se clent ko remove kar do
				delete(h.rooms[room], client)
			}
			// close the channel for this client
			close(client.send)

		case event := <-h.broadcast:
			// Is event se related koi client hai
			if clients, ok := h.rooms[event.RoomId]; ok {
				msg, _ := json.Marshal(event)
				for c := range clients {
					select {
					case c.send <- msg:
					default:
						delete(h.clients, c)
					}
				}
			}
		}
	}
}
