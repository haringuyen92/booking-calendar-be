package staff_requests

type GetAllStaffRequest struct {
	StoreID  uint   `json:"store_id"`
	Page     int    `form:"page" binding:"omitempty,min=1"`
	PageSize int    `form:"page_size" binding:"omitempty,min=1,max=100"`
	SortBy   string `form:"sort_by" binding:"omitempty"`
	SortDesc bool   `form:"sort_desc"`
	Search   string `form:"search"`
}
