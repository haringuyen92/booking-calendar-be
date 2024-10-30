package common_dto

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type SettingTimeDto struct {
	ID                  primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	StoreID             uint               `json:"store_id" bson:"store_id,omitempty"`
	IsOpen              bool               `json:"is_open" bson:"is_open"`
	IsApplyDailySetting bool               `json:"is_apply_daily_setting" bson:"is_apply_daily_setting"`
	DailySetting        SettingTimeSlot    `json:"daily_setting" bson:"daily_setting"`
	MondaySetting       SettingTimeSlot    `json:"monday_setting" bson:"monday_setting"`
	TuesdaySetting      SettingTimeSlot    `json:"tuesday_setting" bson:"tuesday_setting"`
	WednesdaySetting    SettingTimeSlot    `json:"wednesday_setting" bson:"wednesday_setting"`
	ThursdaySetting     SettingTimeSlot    `json:"thursday_setting" bson:"thursday_setting"`
	FridaySetting       SettingTimeSlot    `json:"friday_setting" bson:"friday_setting"`
	SaturdaySetting     SettingTimeSlot    `json:"saturday_setting" bson:"saturday_setting"`
	SundaySetting       SettingTimeSlot    `json:"sunday_setting" bson:"sunday_setting"`
}

type SettingTimeSlot struct {
	IsOpenAllDay bool                  `json:"is_open_all_day" bson:"is_open_all_day"`
	IsOffDay     bool                  `json:"is_off_day" bson:"is_off_day"`
	SlotTime     []SettingTimeSlotItem `json:"slot_time,omitempty" bson:"slot_time,omitempty"`
}

type SettingTimeSlotItem struct {
	TimeStart *time.Time `json:"time_start" bson:"time_start"`
	TimeEnd   *time.Time `json:"time_end" bson:"time_end"`
}
