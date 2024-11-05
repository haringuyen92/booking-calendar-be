package staff_dto

import "booking-calendar-server-backend/internal/core/enums"

type UpdateStaffDto struct {
	Name           string                 `json:"name"`
	Email          string                 `json:"email"`
	Phone          string                 `json:"phone"`
	Cost           uint                   `json:"cost"`
	MaxBookingSlot uint                   `json:"max_booking_slot"`
	Active         enums.StaffActive      `json:"active"`
	Color          string                 `json:"color"`
	Position       uint                   `json:"position"`
	IsAllCourse    enums.StaffIsAllCourse `json:"is_all_course"`
}
