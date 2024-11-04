package store_requests

import store_dto "booking-calendar-server-backend/internal/modules/store/dto"

type UpdateSettingSlotRequest struct {
	StoreID uint                            `json:"store_id"`
	Setting *store_dto.UpdateSettingSlotDto `json:"setting"`
}
