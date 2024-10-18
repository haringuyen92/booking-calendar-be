package booking_dto

import (
	"time"
)

type CreateBookingDto struct {
	User      CreateBookingUserDto  `json:"user" bson:"user"`
	Store     CreateBookingStoreDto `json:"store" bson:"store"`
	StartTime time.Time             `json:"start_time" bson:"start_time"`
	EndTime   time.Time             `json:"end_time" bson:"end_time"`
	CreatedAt *time.Time            `json:"created_at" bson:"created_at"`
	UpdatedAt *time.Time            `json:"updated_at" bson:"updated_at"`
	DeletedAt *time.Time            `json:"deleted_at" bson:"deleted_at"`
}
type CreateBookingUserDto struct {
	UserID uint64 `json:"user_id" bson:"user_id"`
	Name   string `json:"name" bson:"name"`
	Email  string `json:"email" bson:"email"`
}

type CreateBookingStoreDto struct {
	StoreID uint64 `json:"store_id" bson:"store_id"`
	Name    string `json:"name" bson:"name"`
	Address string `json:"address" bson:"address"`
}
