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
