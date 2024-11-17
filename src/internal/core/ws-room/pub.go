package ws_room

type BroadcastRoute string

const (
	OnUserUpdated BroadcastRoute = "OnUserUpdated"
)

type Message struct {
	RoomID string      `json:"room_id"`
	Data   interface{} `json:"data"`
	Type   string      `json:"type"`
}

func Broadcast(room string, message Message) error {
	return nil
}
