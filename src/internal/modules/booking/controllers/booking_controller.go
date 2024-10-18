package booking_controllers

import (
	"booking-calendar-server-backend/internal/core/base/interceptors"
	"booking-calendar-server-backend/internal/modules/booking/dto"
	"booking-calendar-server-backend/internal/modules/booking/filters"
	"booking-calendar-server-backend/internal/modules/booking/requests"
	"booking-calendar-server-backend/internal/modules/booking/services"
	"github.com/gin-gonic/gin"
)

type BookingController struct {
	bookingService booking_services.BookingService
}

func NewBookingController(
	bookingService booking_services.BookingService,
) *BookingController {
	return &BookingController{
		bookingService: bookingService,
	}
}

func (co *BookingController) Create(c *gin.Context, req *booking_requests.CreateBookingRequest) error {
	err := co.bookingService.Create(&booking_dto.CreateBookingDto{})
	if err != nil {
		return interceptors.ResponseError(c, err)
	}
	return interceptors.ResponseSuccess(c, nil)
}

func (co *BookingController) Delete(c *gin.Context, req *booking_requests.DeleteBookingRequest) error {
	err := co.bookingService.Delete(&booking_filters.BookingFilter{})
	if err != nil {
		return interceptors.ResponseError(c, err)
	}
	return interceptors.ResponseSuccess(c, nil)
}

func (co *BookingController) Update(c *gin.Context, req *booking_requests.UpdateBookingRequest) error {
	err := co.bookingService.Update(
		&booking_filters.BookingFilter{ID: req.ID},
		&booking_dto.UpdateBookingDto{},
	)
	if err != nil {
		return interceptors.ResponseError(c, err)
	}
	return interceptors.ResponseSuccess(c, nil)
}

func (co *BookingController) GetOne(c *gin.Context, req *booking_requests.GetBookingRequest) error {
	res, err := co.bookingService.GetOne(&booking_filters.BookingFilter{
		ID: req.ID,
	})

	if err != nil {
		return interceptors.ResponseError(c, err)
	}
	return interceptors.ResponseSuccess(c, res)
}

func (co *BookingController) GetAll(c *gin.Context, req *booking_requests.GetAllBookingRequest) error {
	res, err := co.bookingService.GetMany(&booking_filters.BookingFilter{
		UserID: req.UserID,
	})

	if err != nil {
		return interceptors.ResponseError(c, err)
	}
	return interceptors.ResponseSuccess(c, res)
}
