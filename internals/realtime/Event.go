package realtime

type Event struct {
	Type    string      `json:"type"`
	RoomId  string      `json:"roomId"`
	Payload interface{} `json:"payload"`
}
