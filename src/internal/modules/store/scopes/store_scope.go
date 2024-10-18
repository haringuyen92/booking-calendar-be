package store_scopes

import (
	store_filter "booking-calendar-server-backend/internal/modules/store/filters"
	"gorm.io/gorm"
)

func StoreScopes(filter *store_filter.StoreFilter) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if filter.ID != 0 {
			db.Where("id = ?", filter.ID)
		}
		if filter.UserID != 0 {
			db.Where("user_id = ?", filter.UserID)
		}
		return db
	}
}
