package common_dto

import (
	"booking-calendar-server-backend/internal/core/enums"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SettingBookingDto struct {
	ID                    primitive.ObjectID       `json:"id" bson:"_id,omitempty"`
	StoreID               uint                     `json:"store_id" bson:"store_id,omitempty"`
	Duration              enums.Duration           `json:"duration" bson:"duration,omitempty"`
	MaxBookingTheSameTime uint                     `json:"max_booking_the_same_time" bson:"max_booking_the_same_time,omitempty"`
	Approve               SettingBookingRequestDto `json:"approve" bson:"approve"`
	Change                SettingBookingRequestDto `json:"change" bson:"change"`
	Cancel                SettingBookingRequestDto `json:"cancel" bson:"cancel"`
}

type SettingBookingRequestDto struct {
	Require        bool `json:"is_require" bson:"is_require"`
	WithBeforeTime uint `json:"with_before_time" bson:"with_before_time"`
}
