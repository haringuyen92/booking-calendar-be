package staff_models

import (
	"gorm.io/gorm"
)

type Staff struct {
	gorm.Model
	UserID   uint   `json:"user_id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Position string `json:"position"`
	// Thêm các trường khác nếu cần
}
