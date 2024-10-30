package store_controllers

import (
	"booking-calendar-server-backend/internal/core/base/interceptors"
	store_dto "booking-calendar-server-backend/internal/modules/store/dto"
	store_filter "booking-calendar-server-backend/internal/modules/store/filters"
	store_requests "booking-calendar-server-backend/internal/modules/store/requests"
	store_services "booking-calendar-server-backend/internal/modules/store/services"
	"github.com/gin-gonic/gin"
)

type StoreController struct {
	storeService store_services.StoreService
}

func NewStoreController(
	storeService store_services.StoreService,
) *StoreController {
	return &StoreController{
		storeService: storeService,
	}
}

func (co *StoreController) GetOne(c *gin.Context, req *store_requests.GetStoreRequest) error {
	res, err := co.storeService.GetOne(&store_filter.StoreFilter{
		ID: req.ID,
	})
	if err != nil {
		return interceptors.ResponseError(c, err)
	}
	return interceptors.ResponseSuccess(c, res)
}

func (co *StoreController) GetAll(c *gin.Context, req *store_requests.GetAllStoreRequest) error {
	res, err := co.storeService.GetMany(&store_filter.StoreFilter{
		UserID: req.UserID,
	})
	if err != nil {
		return interceptors.ResponseError(c, err)
	}
	return interceptors.ResponseSuccess(c, res)
}

func (co *StoreController) Create(c *gin.Context, req *store_requests.CreateStoreRequest) error {
	err := co.storeService.Create(&store_dto.CreateStoreDto{
		Name:        req.Name,
		Description: req.Description,
		Address:     req.Address,
		Website:     req.Website,
		Logo:        req.Logo,
		Email:       req.Email,
		Phone:       req.Phone,
		Location:    req.Location,
	})
	if err != nil {
		return interceptors.ResponseError(c, err)
	}
	return interceptors.ResponseSuccess(c, nil)
}

func (co *StoreController) Update(c *gin.Context, req *store_requests.UpdateStoreRequest) error {
	err := co.storeService.Update(
		&store_filter.StoreFilter{
			ID: req.ID,
		},
		&store_dto.UpdateStoreDto{
			Name:        req.Name,
			Description: req.Description,
			Address:     req.Address,
			Website:     req.Website,
			Logo:        req.Logo,
			Email:       req.Email,
			Phone:       req.Phone,
			Location:    req.Location,
			Status:      req.Status,
		},
	)
	if err != nil {
		return interceptors.ResponseError(c, err)
	}
	return interceptors.ResponseSuccess(c, nil)
}

func (co *StoreController) Delete(c *gin.Context, req *store_requests.DeleteStoreRequest) error {
	err := co.storeService.Delete(&store_filter.StoreFilter{
		ID: req.ID,
	})
	if err != nil {
		return interceptors.ResponseError(c, err)
	}
	return interceptors.ResponseSuccess(c, nil)
}
