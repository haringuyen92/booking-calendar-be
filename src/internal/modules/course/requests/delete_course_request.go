package course_requests

type DeleteCourseRequest struct {
	ID      uint `json:"id"`
	StoreID uint `json:"store_id"`
}
