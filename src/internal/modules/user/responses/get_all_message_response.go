package user_responses

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type GetAllMessageResponse struct {
	ID             primitive.ObjectID `json:"id"`
	ConversationID string             `json:"conversation_id"`
	SenderID       uint               `json:"sender_id"`
	Content        string             `json:"content"`
	IsRead         bool               `json:"is_read"`
	CreatedAt      time.Time          `json:"created_at"`
}
