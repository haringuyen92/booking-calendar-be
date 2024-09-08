package base

type Response struct {
	Data       interface{} `json:"data"`
	Code       uint        `json:"code"`
	Message    string      `json:"message"`
	Pagination Pagination  `json:"pagination,omitempty"`
}

type Pagination struct {
	PerPage     uint `json:"per_page,omitempty"`
	CurrentPage uint `json:"current_page,omitempty"`
	TotalCount  uint `json:"total_count,omitempty"`
}
