package mappers

import (
	"booking-calendar-server-backend/internal/modules/booking/models"
	"booking-calendar-server-backend/internal/modules/booking/responses"
)

func GetBookingResponseMapper(model *models.Booking) *responses.GetBookingResponse {
	return &responses.GetBookingResponse{
		ID:        model.ID.Hex(),
		CreateAt:  model.CreatedAt,
		UpdateAt:  model.UpdatedAt,
		StartTime: model.StartTime,
		EndTime:   model.EndTime,
	}
}
