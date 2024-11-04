package store_controllers

import (
	"booking-calendar-server-backend/internal/core/base/interceptors"
	store_filter "booking-calendar-server-backend/internal/modules/store/filters"
	store_formatters "booking-calendar-server-backend/internal/modules/store/formatters"
	"booking-calendar-server-backend/internal/modules/store/requests/settings"
	store_service "booking-calendar-server-backend/internal/modules/store/services/setting"
	"github.com/gin-gonic/gin"
)

type SettingController struct {
	settingStoreService store_service.SettingStoreService
}

func NewSettingController(
	settingStoreService store_service.SettingStoreService,
) *SettingController {
	return &SettingController{
		settingStoreService: settingStoreService,
	}
}

func (co *SettingController) GetSettingTime(c *gin.Context, req *store_requests.GetSettingTimeRequest) error {
	res, err := co.settingStoreService.GetSettingTime(&store_filter.SettingTimeFilter{
		StoreID: req.StoreID,
	})
	if err != nil {
		return interceptors.ResponseError(c, err)
	}
	return interceptors.ResponseSuccess(c, res)
}

func (co *SettingController) UpdateSettingTime(c *gin.Context, req *store_requests.UpdateSettingTimeRequest) error {
	err := co.settingStoreService.UpdateSettingTime(
		req.StoreID,
		store_formatters.UpdateSettingTimeFormatter(req),
	)
	if err != nil {
		return interceptors.ResponseError(c, err)
	}
	return interceptors.ResponseSuccess(c, nil)
}

func (co *SettingController) GetSettingBooking(c *gin.Context, req *store_requests.GetSettingBookingRequest) error {
	res, err := co.settingStoreService.GetSettingBooking(&store_filter.SettingBookingFilter{
		StoreID: req.StoreID,
	})
	if err != nil {
		return interceptors.ResponseError(c, err)
	}
	return interceptors.ResponseSuccess(c, res)
}

func (co *SettingController) UpdateSettingBooking(c *gin.Context, req *store_requests.UpdateSettingBookingRequest) error {
	err := co.settingStoreService.UpdateSettingBooking(
		req.StoreID,
		store_formatters.UpdateSettingBookingFormatter(req),
	)
	if err != nil {
		return interceptors.ResponseError(c, err)
	}
	return interceptors.ResponseSuccess(c, nil)
}

func (co *SettingController) GetSettingSlot(c *gin.Context, req *store_requests.GetSettingSlotRequest) error {
	res, err := co.settingStoreService.GetSettingSlot(&store_filter.SettingSlotFilter{
		StoreID: req.StoreID,
	})
	if err != nil {
		return interceptors.ResponseError(c, err)
	}
	return interceptors.ResponseSuccess(c, res)
}

func (co *SettingController) UpdateSettingSlot(c *gin.Context, req *store_requests.UpdateSettingSlotRequest) error {
	err := co.settingStoreService.UpdateSettingSlot(
		req.StoreID,
		store_formatters.UpdateSettingSlotFormatter(req),
	)
	if err != nil {
		return interceptors.ResponseError(c, err)
	}
	return interceptors.ResponseSuccess(c, nil)
}
