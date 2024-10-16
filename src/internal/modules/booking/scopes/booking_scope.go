package scopes

import (
	"booking-calendar-server-backend/internal/modules/booking/filters"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func BookingScope(filter *filters.BookingFilter) bson.D {
	query := bson.D{}
	if filter.ID != "" {
		mongoID, err := primitive.ObjectIDFromHex(filter.ID)
		if err != nil {
			return query
		}
		query = append(query, bson.E{Key: "_id", Value: mongoID})
	}
	if filter.UserID != 0 {
		query = append(query, bson.E{Key: "user.id", Value: filter.UserID})
	}
	if filter.StoreID != 0 {
		query = append(query, bson.E{Key: "store.id", Value: filter.StoreID})
	}
	if filter.IsDeleted {
		query = append(query, bson.E{Key: "deleted_at", Value: bson.E{Key: "$ne", Value: nil}})
	}
	return query
}
