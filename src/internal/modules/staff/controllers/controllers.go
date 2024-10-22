package staff_controllers

import (
	"booking-calendar-server-backend/internal/core/base/interceptors"
	staff_dto "booking-calendar-server-backend/internal/modules/staff/dto"
	staff_filter "booking-calendar-server-backend/internal/modules/staff/filters"
	staff_requests "booking-calendar-server-backend/internal/modules/staff/requests"
	staff_services "booking-calendar-server-backend/internal/modules/staff/services"

	"github.com/gin-gonic/gin"
)

type StaffController struct {
	staffService staff_services.StaffService
}

func NewStaffController(staffService staff_services.StaffService) *StaffController {
	return &StaffController{
		staffService: staffService,
	}
}

func (co *StaffController) GetAll(c *gin.Context, req *staff_requests.GetAllStaffRequest) error {
	res, err := co.staffService.GetMany(
		&staff_filter.StaffFilter{},
	)
	if err != nil {
		return interceptors.ResponseError(c, err)
	}
	return interceptors.ResponseSuccess(c, res)
}

func (co *StaffController) GetOne(c *gin.Context, req *staff_requests.GetStaffRequest) error {
	res, err := co.staffService.GetOne(
		&staff_filter.StaffFilter{ID: req.ID},
	)
	if err != nil {
		return interceptors.ResponseError(c, err)
	}
	return interceptors.ResponseSuccess(c, res)
}

func (co *StaffController) Create(c *gin.Context, req *staff_requests.CreateStaffRequest) error {
	err := co.staffService.Create(
		&staff_dto.CreateStaffDto{
			Name: req.Name,
		},
	)
	if err != nil {
		return interceptors.ResponseError(c, err)
	}
	return interceptors.ResponseSuccess(c, nil)
}

func (co *StaffController) Update(c *gin.Context, req *staff_requests.UpdateStaffRequest) error {
	err := co.staffService.Update(
		&staff_filter.StaffFilter{ID: req.ID},
		&staff_dto.UpdateStaffDto{
			Name: req.Name,
		},
	)
	if err != nil {
		return interceptors.ResponseError(c, err)
	}
	return interceptors.ResponseSuccess(c, nil)
}

func (co *StaffController) Delete(c *gin.Context, req *staff_requests.DeleteStaffRequest) error {
	err := co.staffService.Delete(
		&staff_filter.StaffFilter{
			ID: req.ID,
		},
	)
	if err != nil {
		return interceptors.ResponseError(c, err)
	}
	return interceptors.ResponseSuccess(c, nil)
}
