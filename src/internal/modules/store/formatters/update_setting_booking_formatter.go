package store_formatters

import (
	common_dto "booking-calendar-server-backend/internal/core/common_dto/store"
	store_dto "booking-calendar-server-backend/internal/modules/store/dto"
	store_requests "booking-calendar-server-backend/internal/modules/store/requests/settings"
)

func UpdateSettingBookingFormatter(req *store_requests.UpdateSettingBookingRequest) *store_dto.UpdateSettingBookingDto {
	return &store_dto.UpdateSettingBookingDto{
		Setting: &common_dto.SettingBookingDto{
			StoreID:               req.StoreID,
			Duration:              req.Duration,
			Approve:               req.Approve,
			Change:                req.Change,
			Cancel:                req.Cancel,
			MaxBookingTheSameTime: req.MaxBookingTheSameTime,
		},
	}
}
