package controllers

import (
	"booking-calendar-server-backend/internal/core/base/interceptors"
	"booking-calendar-server-backend/internal/modules/booking/dto"
	"booking-calendar-server-backend/internal/modules/booking/filters"
	"booking-calendar-server-backend/internal/modules/booking/requests"
	"booking-calendar-server-backend/internal/modules/booking/services"
	"github.com/gin-gonic/gin"
)

type BookingController struct {
	bookingService services.BookingService
}

func NewBookingController(
	bookingService services.BookingService,
) *BookingController {
	return &BookingController{
		bookingService: bookingService,
	}
}

func (co *BookingController) Create(c *gin.Context, req *requests.CreateBookingRequest) error {
	err := co.bookingService.Create(&dto.CreateBookingDto{})
	if err != nil {
		return interceptors.ResponseError(c, err)
	}
	return interceptors.ResponseSuccess(c, nil)
}

func (co *BookingController) Delete(c *gin.Context, req *requests.DeleteBookingRequest) error {
	err := co.bookingService.Delete(&filters.BookingFilter{})
	if err != nil {
		return interceptors.ResponseError(c, err)
	}
	return interceptors.ResponseSuccess(c, nil)
}

func (co *BookingController) Update(c *gin.Context, req *requests.UpdateBookingRequest) error {
	err := co.bookingService.Update(
		&filters.BookingFilter{ID: req.ID},
		&dto.UpdateBookingDto{},
	)
	if err != nil {
		return interceptors.ResponseError(c, err)
	}
	return interceptors.ResponseSuccess(c, nil)
}

func (co *BookingController) Get(c *gin.Context, req *requests.GetBookingRequest) error {
	res, err := co.bookingService.GetOne(&filters.BookingFilter{
		ID: req.ID,
	})

	if err != nil {
		return interceptors.ResponseError(c, err)
	}
	return interceptors.ResponseSuccess(c, res)
}

func (co *BookingController) GetAll(c *gin.Context, req *requests.GetAllBookingRequest) error {
	res, err := co.bookingService.GetMany(&filters.BookingFilter{
		UserID: req.UserID,
	})

	if err != nil {
		return interceptors.ResponseError(c, err)
	}
	return interceptors.ResponseSuccess(c, res)
}
