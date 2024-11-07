package user_scopes

import (
	user_filters "booking-calendar-server-backend/internal/modules/user/filters"
	"go.mongodb.org/mongo-driver/bson"
)

func MessageScope(filter *user_filters.MessageFilter) bson.D {
	query := bson.D{}

	if filter.ConversationID != "" {
		query = append(query, bson.E{Key: "conversation_id", Value: filter.ConversationID})
	}
	return query
}
