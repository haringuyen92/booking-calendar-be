package staff_requests

import "booking-calendar-server-backend/internal/core/enums"

type CreateStaffRequest struct {
	StoreID        uint                   `json:"store_id"`
	Name           string                 `json:"name" binding:"required"`
	Email          string                 `json:"email" binding:"required,email"`
	Phone          string                 `json:"phone" binding:"required"`
	Cost           uint                   `json:"cost" binding:"required"`
	MaxBookingSlot uint                   `json:"max_booking_slot" binding:"required"`
	Color          string                 `json:"color" binding:"required"`
	Position       uint                   `json:"position" binding:"required"`
	IsAllCourse    enums.StaffIsAllCourse `json:"is_all_course"`
}
