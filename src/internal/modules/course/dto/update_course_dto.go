package course_dto

type UpdateCourseDto struct {
	Name string `json:"name" binding:"required"`
}
