package staff_dto

type UpdateStaffDto struct {
	Name     string `json:"name,omitempty"`
	Email    string `json:"email,omitempty"`
	Position string `json:"position,omitempty"`
	// Thêm các trường khác nếu cần
}
