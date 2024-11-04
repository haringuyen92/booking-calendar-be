package store_dto

import common_dto "booking-calendar-server-backend/internal/core/common_dto/store"

type UpdateSettingSlotDto struct {
	Setting *common_dto.SettingSlotDto `json:"setting"`
}
