package store_service

import (
	common_dto "booking-calendar-server-backend/internal/core/common_dto/store"
	store_dto "booking-calendar-server-backend/internal/modules/store/dto"
	store_filter "booking-calendar-server-backend/internal/modules/store/filters"
)

func (s *settingStoreService) GetSettingTime(filter *store_filter.SettingTimeFilter) (*common_dto.SettingTimeDto, error) {
	return s.settingTimeRepository.GetSettingTime(filter.StoreID)
}

func (s *settingStoreService) UpdateSettingTime(storeID uint, dto *store_dto.UpdateSettingTimeDto) error {
	return s.settingTimeRepository.UpdateSettingTime(
		storeID,
		&common_dto.SettingTimeDto{
			IsOpen:              dto.IsOpen,
			IsApplyDailySetting: dto.IsApplyDailySetting,
			DailySetting:        dto.DailySetting,
			MondaySetting:       dto.MondaySetting,
			TuesdaySetting:      dto.TuesdaySetting,
			WednesdaySetting:    dto.WednesdaySetting,
			ThursdaySetting:     dto.ThursdaySetting,
			FridaySetting:       dto.FridaySetting,
			SaturdaySetting:     dto.SaturdaySetting,
			SundaySetting:       dto.SundaySetting,
		},
	)
}
