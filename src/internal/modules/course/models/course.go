package course_models

import (
	"booking-calendar-server-backend/internal/core/enums"
	"gorm.io/gorm"
)

type Course struct {
	gorm.Model
	StoreID      uint               `json:"store_id"`
	Name         string             `json:"name" gorm:"varchar(50)"`
	Image        string             `json:"image" gorm:"varchar(255)"`
	Description  string             `json:"description"`
	Cost         uint               `json:"cost"`
	EstimateTime uint               `json:"estimate_time"`
	Active       enums.CourseActive `json:"active"`
	Color        string             `json:"color" gorm:"varchar(50)"`
	Position     uint               `json:"position"`
}
