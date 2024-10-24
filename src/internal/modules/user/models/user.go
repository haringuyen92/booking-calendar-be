package user_models

import "booking-calendar-server-backend/internal/core/base"

type User struct {
	base.Model
	Name    string `json:"name" gorm:"type:varchar(255)"`
	Email   string `json:"email" gorm:"type:varchar(255)"`
	Phone   string `json:"phone" gorm:"type:varchar(255)"`
	Address string `json:"address" gorm:"type:varchar(255)"`
	Avatar  string `json:"avatar" gorm:"type:varchar(255)"`
}
