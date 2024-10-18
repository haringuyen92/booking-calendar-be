package booking_mappers

import (
	booking_models "booking-calendar-server-backend/internal/modules/booking/models"
	"booking-calendar-server-backend/internal/modules/booking/responses"
)

func GetAllBookingResponseMapper(bookings []*booking_models.Booking) []*booking_responses.GetAllBookingResponse {
	var response []*booking_responses.GetAllBookingResponse
	for _, booking := range bookings {
		response = append(response, &booking_responses.GetAllBookingResponse{
			ID: booking.ID.Hex(),
		})
	}
	return response
}
