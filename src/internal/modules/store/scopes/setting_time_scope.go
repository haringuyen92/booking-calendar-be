package store_scopes

import (
	store_filter "booking-calendar-server-backend/internal/modules/store/filters"
	"go.mongodb.org/mongo-driver/bson"
)

func SettingTimeScope(filter *store_filter.SettingTimeFilter) bson.D {
	query := bson.D{}
	if filter.StoreID != 0 {
		query = append(query, bson.E{Key: "store_id", Value: filter.StoreID})
	}
	return query
}
