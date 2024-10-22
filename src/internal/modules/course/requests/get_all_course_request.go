package course_requests

type GetAllCourseRequest struct {
	StoreID  uint `json:"store_id"`
	Page     int  `json:"page"`
	PageSize int  `json:"page_size"`
}
