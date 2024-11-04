package store_requests

import (
	store_dto "booking-calendar-server-backend/internal/modules/store/dto"
)

type UpdateSettingBookingRequest struct {
	StoreID uint                               `json:"store_id"`
	Setting *store_dto.UpdateSettingBookingDto `json:"setting"`
}
