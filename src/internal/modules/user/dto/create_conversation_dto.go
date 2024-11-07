package user_dto

type CreateConversationDto struct {
	SenderID   uint `json:"sender_id"`
	ReceiverID uint `json:"receiver_id"`
}
