package user_requests

type CreateMessageRequest struct {
	ConversationID string `json:"conversation_id"`
	Content        string `json:"content"`
}
