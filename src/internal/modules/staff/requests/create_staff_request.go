package staff_requests

import "booking-calendar-server-backend/internal/core/enums"

type CreateStaffRequest struct {
	StoreID        uint                   `json:"store_id"`
	Name           string                 `json:"name"`
	Email          string                 `json:"email"`
	Phone          string                 `json:"phone"`
	Cost           uint                   `json:"cost"`
	MaxBookingSlot uint                   `json:"max_booking_slot"`
	Color          string                 `json:"color"`
	Position       uint                   `json:"position"`
	IsAllCourse    enums.StaffIsAllCourse `json:"is_all_course"`
}
