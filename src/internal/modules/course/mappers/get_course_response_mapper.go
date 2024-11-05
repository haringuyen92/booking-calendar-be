package course_mappers

import (
	course_models "booking-calendar-server-backend/internal/modules/course/models"
	course_responses "booking-calendar-server-backend/internal/modules/course/responses"
)

func GetCourseResponseMapper(course *course_models.Course) *course_responses.GetCourseResponse {
	return &course_responses.GetCourseResponse{
		ID:           course.ID,
		StoreID:      course.StoreID,
		Name:         course.Name,
		Image:        course.Image,
		Description:  course.Description,
		Cost:         course.Cost,
		Active:       course.Active,
		EstimateTime: course.EstimateTime,
		Position:     course.Position,
		CreatedAt:    course.CreatedAt,
		UpdatedAt:    course.UpdatedAt,
	}
}
