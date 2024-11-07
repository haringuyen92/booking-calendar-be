package user_models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Conversation struct {
	ID           primitive.ObjectID        `bson:"_id,omitempty" json:"id,omitempty"`
	Participants map[uint]ConversationUser `bson:"participants" json:"participants"`
	CreatedAt    time.Time                 `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time                 `bson:"updated_at" json:"updated_at"`
	DeletedAt    *time.Time                `bson:"deleted_at,omitempty" json:"-"`
	LastMessage  *ConversationLastMessage  `bson:"last_message,omitempty" json:"last_message,omitempty"`
}

type ConversationLastMessage struct {
	SenderID uint      `bson:"sender_id" json:"sender_id"`
	Content  string    `bson:"content" json:"content"`
	SendAt   time.Time `bson:"send_at" json:"send_at"`
}

type ConversationUser struct {
	Name   string `bson:"name" json:"name"`
	Email  string `bson:"email" json:"email"`
	Avatar string `bson:"avatar" json:"avatar"`
}
