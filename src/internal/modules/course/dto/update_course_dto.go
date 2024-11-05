package course_dto

import "booking-calendar-server-backend/internal/core/enums"

type UpdateCourseDto struct {
	StoreID      uint               `json:"store_id"`
	Name         string             `json:"name"`
	Image        string             `json:"image"`
	Description  string             `json:"description"`
	Cost         uint               `json:"cost"`
	Active       enums.CourseActive `json:"active"`
	EstimateTime uint               `json:"estimate_time"`
	Position     uint               `json:"position"`
}
