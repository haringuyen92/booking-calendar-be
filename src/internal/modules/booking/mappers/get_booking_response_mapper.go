package mappers

import (
	"booking-calendar-server-backend/internal/modules/booking/models"
	"booking-calendar-server-backend/internal/modules/booking/responses"
)

func GetBookingResponseMapper(model *models.Booking) *responses.GetBookingResponse {
	return &responses.GetBookingResponse{}
}
