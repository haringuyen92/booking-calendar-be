package user_requests

type CreateConversationRequest struct {
	SenderID   uint `json:"sender_id"`
	ReceiverID uint `json:"receiver_id"`
}
