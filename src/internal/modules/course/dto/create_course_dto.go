package course_dto

type CreateCourseDto struct {
	Name string `json:"name" binding:"required"`
}
