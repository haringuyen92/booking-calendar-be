package staff_filter

type StaffFilter struct {
	ID      uint   `json:"id,omitempty"`
	StoreID uint   `json:"store_id,omitempty"`
	UserID  uint   `json:"user_id,omitempty"`
	Name    string `json:"name,omitempty"`
	Email   string `json:"email,omitempty"`
	// Thêm các trường khác nếu cần
}
