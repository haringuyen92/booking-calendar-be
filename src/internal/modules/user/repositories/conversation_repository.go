package user_repositories

import (
	user_filters "booking-calendar-server-backend/internal/modules/user/filters"
	user_models "booking-calendar-server-backend/internal/modules/user/models"
	user_scopes "booking-calendar-server-backend/internal/modules/user/scopes"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type ConversationRepository interface {
	GetByUserID(userID uint) ([]*user_models.Conversation, error)
	Create(model *user_models.Conversation) error
	UpdateLastMessage(ID string, lastMessage *user_models.ConversationLastMessage) error
	Update(filter *user_filters.ConversationFilter, update *user_models.Conversation) error
	GetMany(filter *user_filters.ConversationFilter) ([]*user_models.Conversation, error)
}

func NewConversationRepository(
	mongoDB *mongo.Database,
) ConversationRepository {
	collection := mongoDB.Collection("conversations")
	return &conversationRepository{
		collection: collection,
	}
}

type conversationRepository struct {
	collection *mongo.Collection
}

func (r *conversationRepository) GetByUserID(userID uint) ([]*user_models.Conversation, error) {
	return r.GetMany(&user_filters.ConversationFilter{
		UserID: userID,
	})
}

func (r *conversationRepository) Create(model *user_models.Conversation) error {
	now := time.Now()
	model.CreatedAt = now
	model.UpdatedAt = now
	model.DeletedAt = nil
	_, err := r.collection.InsertOne(context.Background(), model)
	if err != nil {
		return err
	}
	return nil
}

func (r *conversationRepository) UpdateLastMessage(ID string, lastMessage *user_models.ConversationLastMessage) error {
	conversation := &user_models.Conversation{
		LastMessage: lastMessage,
	}
	return r.Update(&user_filters.ConversationFilter{
		ID: ID,
	}, conversation)
}

func (r *conversationRepository) Update(filter *user_filters.ConversationFilter, update *user_models.Conversation) error {
	scope := user_scopes.ConversationScope(filter)
	update.UpdatedAt = time.Now()
	_, err := r.collection.UpdateOne(
		context.Background(),
		scope,
		bson.M{"$set": update},
	)
	if err != nil {
		return err
	}
	return nil
}

func (r *conversationRepository) GetMany(filter *user_filters.ConversationFilter) ([]*user_models.Conversation, error) {
	var conversations []*user_models.Conversation
	cursor, err := r.collection.Find(context.Background(), user_scopes.ConversationScope(filter))
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var conversation user_models.Conversation
		if err := cursor.Decode(&conversation); err != nil {
			return nil, err
		}
		conversations = append(conversations, &conversation)
	}

	return conversations, nil
}
