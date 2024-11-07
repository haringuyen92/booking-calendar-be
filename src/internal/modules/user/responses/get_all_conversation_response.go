package user_responses

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type GetAllConversationsResponse struct {
	ID           primitive.ObjectID      `json:"id,omitempty"`
	Participants []ConversationUser      `json:"participants"`
	CreatedAt    time.Time               `json:"created_at"`
	UpdatedAt    time.Time               `json:"updated_at"`
	LastMessage  ConversationLastMessage `json:"last_message"`
}

type ConversationLastMessage struct {
	SenderID uint      `json:"sender_id"`
	Content  string    `json:"content"`
	SendAt   time.Time `json:"send_at"`
}

type ConversationUser struct {
	UserID uint   `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Avatar string `json:"avatar"`
}
