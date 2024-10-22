package staff_requests

type CreateStaffRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Phone    string `json:"phone" binding:"required"`
	Position string `json:"position" binding:"required"`
}
