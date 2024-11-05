package course_responses

import (
	"booking-calendar-server-backend/internal/core/enums"
	"time"
)

type GetAllCourseResponse struct {
	ID           uint               `json:"id"`
	StoreID      uint               `json:"store_id"`
	Name         string             `json:"name"`
	Image        string             `json:"image"`
	Description  string             `json:"description"`
	Cost         uint               `json:"cost"`
	Active       enums.CourseActive `json:"active"`
	EstimateTime uint               `json:"estimate_time"`
	Position     uint               `json:"position"`
	CreatedAt    time.Time          `json:"created_at"`
	UpdatedAt    time.Time          `json:"updated_at"`
}
