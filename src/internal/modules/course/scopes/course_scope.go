package course_scopes

import (
	course_filters "booking-calendar-server-backend/internal/modules/course/filters"

	"gorm.io/gorm"
)

func CourseScope(filter *course_filters.CourseFilter) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if filter.ID != 0 {
			db = db.Where("id = ?", filter.ID)
		}

		return db
	}
}
