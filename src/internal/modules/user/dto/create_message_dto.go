package user_dto

type CreateMessageDto struct {
	SenderID       uint        `json:"sender_id"`
	ConversationID string      `json:"conversation_id"`
	Content        interface{} `json:"content"`
}
