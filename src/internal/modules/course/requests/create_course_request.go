package course_requests

type CreateCourseRequest struct {
	StoreID      uint   `json:"store_id"`
	Name         string `json:"name"`
	Image        string `json:"image"`
	Description  string `json:"description"`
	Cost         uint   `json:"cost"`
	EstimateTime uint   `json:"estimate_time"`
	Position     uint   `json:"position"`
}
