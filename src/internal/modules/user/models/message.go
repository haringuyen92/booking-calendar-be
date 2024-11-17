package user_models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Message struct {
	ID             primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ConversationID string             `json:"conversation_id" bson:"conversation_id,omitempty"`
	SenderID       uint               `json:"sender_id" bson:"sender_id,omitempty"`
	Content        string             `json:"content" bson:"content,omitempty"`
	IsRead         bool               `json:"is_read" bson:"is_read"`
	CreatedAt      time.Time          `json:"created_at" bson:"created_at,omitempty"`
	UpdatedAt      time.Time          `json:"updated_at" bson:"updated_at,omitempty"`
	DeletedAt      time.Time          `json:"deleted_at" bson:"deleted_at,omitempty"`
}
