package course_requests

type UpdateCourseRequest struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
