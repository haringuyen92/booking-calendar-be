package store_requests

import common_dto "booking-calendar-server-backend/internal/core/common_dto/store"

type UpdateSettingSlotRequest struct {
	StoreID                   uint                             `json:"store_id"`
	Course                    common_dto.SettingSlotRequestDto `json:"course"`
	Staff                     common_dto.SettingSlotRequestDto `json:"staff"`
	DefaultCourseEstimateTime uint                             `json:"default_course_estimate_time"`
}
