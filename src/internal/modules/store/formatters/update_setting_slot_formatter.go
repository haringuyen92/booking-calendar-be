package store_formatters

import (
	common_dto "booking-calendar-server-backend/internal/core/common_dto/store"
	store_dto "booking-calendar-server-backend/internal/modules/store/dto"
	store_requests "booking-calendar-server-backend/internal/modules/store/requests/settings"
)

func UpdateSettingSlotFormatter(req *store_requests.UpdateSettingSlotRequest) *store_dto.UpdateSettingSlotDto {
	return &store_dto.UpdateSettingSlotDto{
		Setting: &common_dto.SettingSlotDto{
			StoreID:                   req.StoreID,
			Course:                    req.Course,
			Staff:                     req.Staff,
			DefaultCourseEstimateTime: req.DefaultCourseEstimateTime,
		},
	}
}
