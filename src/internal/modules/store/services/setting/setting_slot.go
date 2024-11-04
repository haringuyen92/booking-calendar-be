package store_service

import (
	common_dto "booking-calendar-server-backend/internal/core/common_dto/store"
	store_dto "booking-calendar-server-backend/internal/modules/store/dto"
	store_filter "booking-calendar-server-backend/internal/modules/store/filters"
)

func (s *settingStoreService) GetSettingSlot(filter *store_filter.SettingSlotFilter) (*common_dto.SettingSlotDto, error) {
	return s.settingSlotRepository.GetSettingByStoreID(filter.StoreID)
}

func (s *settingStoreService) UpdateSettingSlot(storeID uint, dto *store_dto.UpdateSettingSlotDto) error {
	return s.settingSlotRepository.UpdateSettingByStoreID(storeID, dto.Setting)
}
