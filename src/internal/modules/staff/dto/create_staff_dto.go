package staff_dto

type CreateStaffDto struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Position string `json:"position"`
	// Thêm các trường khác nếu cần
}
