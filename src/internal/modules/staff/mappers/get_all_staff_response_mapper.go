package staff_mappers

import (
	staff_models "booking-calendar-server-backend/internal/modules/staff/models"
	staff_responses "booking-calendar-server-backend/internal/modules/staff/responses"
)

func GetAllStaffResponseMapper(staffs []*staff_models.Staff) []*staff_responses.GetAllStaffResponse {
	var staffResponses []*staff_responses.GetAllStaffResponse
	for _, staff := range staffs {
		staffResponses = append(staffResponses, &staff_responses.GetAllStaffResponse{
			ID:             staff.ID,
			StoreID:        staff.StoreID,
			Name:           staff.Name,
			Email:          staff.Email,
			Phone:          staff.Phone,
			Cost:           staff.Cost,
			MaxBookingSlot: staff.MaxBookingSlot,
			Active:         staff.Active,
			Color:          staff.Color,
			Position:       staff.Position,
			IsAllCourse:    staff.IsAllCourse,
		})
	}
	return staffResponses
}
