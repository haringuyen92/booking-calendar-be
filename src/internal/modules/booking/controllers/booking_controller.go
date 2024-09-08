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

func (co *BookingController) CreateBooking(c *gin.Context, req *requests.CreateBookingRequest) error {
	err := co.bookingService.CreateBooking(&dto.CreateBookingDto{})
	if err != nil {
		return interceptors.ResponseError(c, err)
	}
	return interceptors.ResponseSuccess(c, nil)
}

func (co *BookingController) DeleteBooking(c *gin.Context, req *requests.DeleteBookingRequest) error {
	err := co.bookingService.DeleteBooking(&filters.BookingFilter{})
	if err != nil {
		return interceptors.ResponseError(c, err)
	}
	return interceptors.ResponseSuccess(c, nil)
}

func (co *BookingController) UpdateBooking(c *gin.Context, req *requests.UpdateBookingRequest) error {
	err := co.bookingService.UpdateBooking(
		&filters.BookingFilter{ID: req.ID},
		&dto.UpdateBookingDto{},
	)
	if err != nil {
		return interceptors.ResponseError(c, err)
	}
	return interceptors.ResponseSuccess(c, nil)
}

func (co *BookingController) GetBooking(c *gin.Context, req *requests.GetBookingRequest) error {
	res, err := co.bookingService.GetBooking(&filters.BookingFilter{
		ID: req.ID,
	})

	if err != nil {
		return interceptors.ResponseError(c, err)
	}
	return interceptors.ResponseSuccess(c, res)
}

func (co *BookingController) GetBookings(c *gin.Context, req *requests.GetAllBookingRequest) error {
	res, err := co.bookingService.GetBookingList(&filters.BookingFilter{
		UserID: req.UserID,
	})

	if err != nil {
		return interceptors.ResponseError(c, err)
	}
	return interceptors.ResponseSuccess(c, res)
}
