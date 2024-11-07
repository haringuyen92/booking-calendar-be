package user_scopes

import (
	user_filters "booking-calendar-server-backend/internal/modules/user/filters"
	"gorm.io/gorm"
)

func UserScopes(filter *user_filters.UserFilter) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if filter.ID != 0 {
			db.Where("id = ?", filter.ID)
		}
		if filter.Name != "" {
			db.Where("name = ?", filter.Name)
		}
		if filter.Email != "" {
			db.Where("email = ?", filter.Email)
		}
		return db
	}
}
