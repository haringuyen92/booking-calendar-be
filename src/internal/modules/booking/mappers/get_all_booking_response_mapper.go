package mappers

import (
	"booking-calendar-server-backend/internal/modules/booking/models"
	"booking-calendar-server-backend/internal/modules/booking/responses"
)

func GetAllBookingResponseMapper(bookings []*models.Booking) []*responses.GetAllBookingResponse {
	var response []*responses.GetAllBookingResponse
	for _, booking := range bookings {
		response = append(response, &responses.GetAllBookingResponse{
			ID: booking.ID,
		})
	}
	return response
}
