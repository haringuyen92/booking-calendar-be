package store_service

import (
	common_dto "booking-calendar-server-backend/internal/core/common_dto/store"
	store_dto "booking-calendar-server-backend/internal/modules/store/dto"
	store_filter "booking-calendar-server-backend/internal/modules/store/filters"
	stores_repositories "booking-calendar-server-backend/internal/modules/store/repositories"
)

type SettingStoreService interface {
	GetSettingTime(filter *store_filter.SettingTimeFilter) (*common_dto.SettingTimeDto, error)
	UpdateSettingTime(storeID uint, dto *store_dto.UpdateSettingTimeDto) error
	GetSettingBooking(filter *store_filter.SettingBookingFilter) (*common_dto.SettingBookingDto, error)
	UpdateSettingBooking(storeID uint, dto *store_dto.UpdateSettingBookingDto) error
}

func NewSettingStoreService(
	settingTimeRepository stores_repositories.SettingTimeRepository,
	settingBookingRepository stores_repositories.SettingBookingRepository,
) SettingStoreService {
	return &settingStoreService{
		settingTimeRepository:    settingTimeRepository,
		settingBookingRepository: settingBookingRepository,
	}
}

type settingStoreService struct {
	settingTimeRepository    stores_repositories.SettingTimeRepository
	settingBookingRepository stores_repositories.SettingBookingRepository
}
