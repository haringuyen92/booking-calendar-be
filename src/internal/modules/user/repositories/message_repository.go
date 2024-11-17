package user_repositories

import (
	user_filters "booking-calendar-server-backend/internal/modules/user/filters"
	user_models "booking-calendar-server-backend/internal/modules/user/models"
	user_scopes "booking-calendar-server-backend/internal/modules/user/scopes"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type MessageRepository interface {
	GetByConversationID(conversationID string) ([]*user_models.Message, error)
	Create(message *user_models.Message) error
	GetMany(filter *user_filters.MessageFilter) ([]*user_models.Message, error)
}

func NewMessageRepository(
	mongoDB *mongo.Database,
) MessageRepository {
	collection := mongoDB.Collection("messages")
	return &messageRepository{
		collection: collection,
	}
}

type messageRepository struct {
	collection *mongo.Collection
}

func (r *messageRepository) GetByConversationID(conversationID string) ([]*user_models.Message, error) {
	return r.GetMany(&user_filters.MessageFilter{
		ConversationID: conversationID,
	})
}

func (r *messageRepository) Create(message *user_models.Message) error {
	now := time.Now()
	message.CreatedAt = now
	message.UpdatedAt = now
	_, err := r.collection.InsertOne(context.Background(), message)
	if err != nil {
		return err
	}
	return nil
}

func (r *messageRepository) GetMany(filter *user_filters.MessageFilter) ([]*user_models.Message, error) {
	var messages []*user_models.Message
	cursor, err := r.collection.Find(context.Background(), user_scopes.MessageScope(filter))
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var message user_models.Message
		if err := cursor.Decode(&message); err != nil {
			return nil, err
		}
		messages = append(messages, &message)
	}
	return messages, nil
}
