package staff_requests

type DeleteStaffRequest struct {
	ID uint `json:"id" binding:"required"`
}
