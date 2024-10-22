package staff_requests

type UpdateStaffRequest struct {
	ID       uint   `json:"id" binding:"required"`
	Name     string `json:"name"`
	Email    string `json:"email" binding:"omitempty,email"`
	Phone    string `json:"phone"`
	Position string `json:"position"`
}
