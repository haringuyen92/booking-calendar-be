package course_mappers

import (
	course_models "booking-calendar-server-backend/internal/modules/course/models"
	course_responses "booking-calendar-server-backend/internal/modules/course/responses"
)

func GetCourseResponseMapper(course *course_models.Course) *course_responses.GetCourseResponse {
	return &course_responses.GetCourseResponse{
		ID:        course.ID,
		Name:      course.Name,
		CreatedAt: course.CreatedAt,
		UpdatedAt: course.UpdatedAt,
	}
}
