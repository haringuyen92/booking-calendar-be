package models

import (
	"booking-calendar-server-backend/internal/core/base"
	"booking-calendar-server-backend/internal/core/enums"
)

type Store struct {
	base.Model
	UserID      uint              `json:"user_id"` // 1-n in user model
	Name        string            `json:"name" gorm:"type:varchar(255)"`
	Description string            `json:"description" gorm:"type:text"`
	Address     string            `json:"address" gorm:"type:varchar(255)"`
	Website     string            `json:"website" gorm:"type:varchar(255)"`
	Logo        string            `json:"logo" gorm:"type:varchar(255)"`
	Email       string            `json:"email" gorm:"type:varchar(255)"`
	Phone       string            `json:"phone" gorm:"type:varchar(255)"`
	Location    string            `json:"location" gorm:"type:varchar(255)"`
	Status      enums.StoreStatus `json:"status"`
}
