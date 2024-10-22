package staff_mappers

import (
	staff_models "booking-calendar-server-backend/internal/modules/staff/models"
	staff_responses "booking-calendar-server-backend/internal/modules/staff/responses"
)

func GetStaffResponseMapper(staff *staff_models.Staff) *staff_responses.GetStaffResponse {
	return &staff_responses.GetStaffResponse{
		ID: staff.ID,
	}
}
