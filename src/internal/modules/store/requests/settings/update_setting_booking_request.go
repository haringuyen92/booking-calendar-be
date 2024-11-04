package store_requests

import (
	common_dto "booking-calendar-server-backend/internal/core/common_dto/store"
	"booking-calendar-server-backend/internal/core/enums"
)

type UpdateSettingBookingRequest struct {
	StoreID               uint                                `json:"store_id"`
	Duration              enums.Duration                      `json:"duration"`
	MaxBookingTheSameTime uint                                `json:"max_booking_the_same_time"`
	Approve               common_dto.SettingBookingRequestDto `json:"approve"`
	Change                common_dto.SettingBookingRequestDto `json:"change"`
	Cancel                common_dto.SettingBookingRequestDto `json:"cancel"`
}
