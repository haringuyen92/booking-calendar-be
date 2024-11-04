package store_service

import (
	common_dto "booking-calendar-server-backend/internal/core/common_dto/store"
	store_dto "booking-calendar-server-backend/internal/modules/store/dto"
	store_filter "booking-calendar-server-backend/internal/modules/store/filters"
	"fmt"
)

func (s *settingStoreService) GetSettingBooking(filter *store_filter.SettingBookingFilter) (*common_dto.SettingBookingDto, error) {
	return s.settingBookingRepository.GetSettingByStoreID(filter.StoreID)
}

func (s *settingStoreService) UpdateSettingBooking(storeID uint, dto *store_dto.UpdateSettingBookingDto) error {
	fmt.Println(storeID)
	return s.settingBookingRepository.UpdateSettingByStoreID(storeID, dto.Setting)
}
