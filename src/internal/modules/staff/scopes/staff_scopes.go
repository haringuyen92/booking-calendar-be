package staff_scopes

import (
	staff_filter "booking-calendar-server-backend/internal/modules/staff/filters"

	"gorm.io/gorm"
)

func StaffScopes(filter *staff_filter.StaffFilter) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if filter.ID != 0 {
			db = db.Where("id = ?", filter.ID)
		}
		if filter.UserID != 0 {
			db = db.Where("user_id = ?", filter.UserID)
		}
		if filter.Name != "" {
			db = db.Where("name LIKE ?", "%"+filter.Name+"%")
		}
		if filter.Email != "" {
			db = db.Where("email = ?", filter.Email)
		}
		// Thêm các điều kiện lọc khác nếu cần
		return db
	}
}
