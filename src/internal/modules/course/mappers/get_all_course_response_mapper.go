package course_mappers

import (
	course_models "booking-calendar-server-backend/internal/modules/course/models"
	course_responses "booking-calendar-server-backend/internal/modules/course/responses"
)

func GetAllCourseResponseMapper(courses []*course_models.Course) []*course_responses.GetAllCourseResponse {
	var response []*course_responses.GetAllCourseResponse

	for _, course := range courses {
		response = append(response, &course_responses.GetAllCourseResponse{
			ID:        course.ID,
			Name:      course.Name,
			CreatedAt: course.CreatedAt,
			UpdatedAt: course.UpdatedAt,
		})
	}

	return response
}
