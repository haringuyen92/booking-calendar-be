package user_scopes

import (
	user_filters "booking-calendar-server-backend/internal/modules/user/filters"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func ConversationScope(filter *user_filters.ConversationFilter) bson.D {
	query := bson.D{}
	if filter.ID != "" {
		mongoID, err := primitive.ObjectIDFromHex(filter.ID)
		if err != nil {
			return nil
		}
		query = append(query, bson.E{Key: "_id", Value: mongoID})
	}
	if filter.UserID != 0 {
		query = append(query, bson.E{
			Key:   fmt.Sprintf("participants.%d", filter.UserID),
			Value: bson.M{"$exists": true},
		})
	}
	query = append(query, bson.E{
		Key:   "deleted_at",
		Value: bson.M{"$exists": false},
	})
	return query
}
