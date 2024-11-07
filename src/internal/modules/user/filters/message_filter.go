package user_filters

type MessageFilter struct {
	UserID         uint   `json:"user_id"`
	ConversationID string `json:"conversation_id"`
}
