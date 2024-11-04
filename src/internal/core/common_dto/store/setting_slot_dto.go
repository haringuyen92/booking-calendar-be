package common_dto

import "go.mongodb.org/mongo-driver/bson/primitive"

type SettingSlotDto struct {
	ID                        primitive.ObjectID    `json:"id" bson:"_id,omitempty"`
	StoreID                   uint                  `json:"store_id" bson:"store_id"`
	Course                    SettingSlotRequestDto `json:"course" bson:"course"`
	Staff                     SettingSlotRequestDto `json:"staff" bson:"staff"`
	DefaultCourseEstimateTime uint                  `json:"default_course_estimate_time" bson:"default_course_estimate_time"`
}

type SettingSlotRequestDto struct {
	Require bool `json:"require" bson:"require"`
	UseCost bool `json:"use_cost" bson:"use_cost"`
}
