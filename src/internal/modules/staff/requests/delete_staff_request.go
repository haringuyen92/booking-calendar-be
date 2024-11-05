package staff_requests

type DeleteStaffRequest struct {
	ID      uint `json:"id" binding:"required"`
	StoreID uint `json:"store_id" binding:"required"`
}
