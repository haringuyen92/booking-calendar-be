package course_requests

type GetCourseRequest struct {
	ID      uint `json:"id"`
	StoreID uint `json:"store_id"`
}
