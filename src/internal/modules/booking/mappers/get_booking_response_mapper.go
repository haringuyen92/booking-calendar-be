package booking_mappers

import (
	booking_models "booking-calendar-server-backend/internal/modules/booking/models"
	booking_responses "booking-calendar-server-backend/internal/modules/booking/responses"
)

func GetBookingResponseMapper(model *booking_models.Booking) *booking_responses.GetBookingResponse {
	return &booking_responses.GetBookingResponse{
		ID:        model.ID.Hex(),
		CreateAt:  model.CreatedAt,
		UpdateAt:  model.UpdatedAt,
		StartTime: model.StartTime,
		EndTime:   model.EndTime,
	}
}
