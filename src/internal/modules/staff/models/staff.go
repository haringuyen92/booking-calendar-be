package staff_models

import (
	"booking-calendar-server-backend/internal/core/enums"
	"gorm.io/gorm"
)

type Staff struct {
	gorm.Model
	StoreID        uint                   `json:"store_id"`
	Name           string                 `json:"name" gorm:"varchar(50)"`
	Email          string                 `json:"email" gorm:"varchar(50)"`
	Phone          string                 `json:"phone" gorm:"varchar(50)"`
	Cost           uint                   `json:"cost"`
	MaxBookingSlot uint                   `json:"max_booking_slot"`
	Active         enums.StaffActive      `json:"active"`
	Color          string                 `json:"color" gorm:"varchar(10)"`
	Position       uint                   `json:"position"`
	IsAllCourse    enums.StaffIsAllCourse `json:"is_all_course"`
}
