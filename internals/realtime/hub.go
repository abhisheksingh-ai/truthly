package realtime

import "encoding/json"

type Hub struct {
	// clients
	Clients map[*Client]bool
	// rooms
	RoomsHub map[string]map[*Client]bool

	// register
	Register chan *Client

	// unregister
	Unregister chan *Client

	//broadcast channer
	Broadcast chan Event
}

func (h *Hub) Run() {
	for {
		select {

		case client := <-h.Register:
			h.Clients[client] = true

		case client := <-h.Unregister:
			// clients se hata do
			delete(h.Clients, client)
			for room := range client.Rooms {
				// Jis bhi image id se ye client connected h vha se clent ko remove kar do
				delete(h.RoomsHub[room], client)
			}
			// close the channel for this client
			close(client.Send)

		case event := <-h.Broadcast:
			// Is event se related koi client hai
			if clients, ok := h.RoomsHub[event.RoomId]; ok {
				msg, _ := json.Marshal(event)
				for c := range clients {
					select {
					case c.Send <- msg:
					default:
						delete(h.Clients, c)
					}
				}
			}
		}
	}
}
