package store_requests

import "time"

type UpdateSettingTimeRequest struct {
	StoreID             uint            `json:"store_id"`
	IsOpen              bool            `json:"is_open"`
	IsApplyDailySetting bool            `json:"is_apply_daily_setting"`
	DailySetting        SettingTimeSlot `json:"daily_setting"`
	MondaySetting       SettingTimeSlot `json:"monday_setting"`
	TuesdaySetting      SettingTimeSlot `json:"tuesday_setting"`
	WednesdaySetting    SettingTimeSlot `json:"wednesday_setting"`
	ThursdaySetting     SettingTimeSlot `json:"thursday_setting"`
	FridaySetting       SettingTimeSlot `json:"friday_setting"`
	SaturdaySetting     SettingTimeSlot `json:"saturday_setting"`
	SundaySetting       SettingTimeSlot `json:"sunday_setting"`
}

type SettingTimeSlot struct {
	IsOpenAllDay bool                  `json:"is_open_all_day"`
	IsOffDay     bool                  `json:"is_off_day"`
	SlotTime     []SettingTimeSlotItem `json:"slot_time,omitempty"`
}

type SettingTimeSlotItem struct {
	TimeStart string `json:"time_start"`
	TimeEnd   string `json:"time_end"`
}

func (s *SettingTimeSlotItem) ConvertToTime() (*time.Time, *time.Time) {
	return convertToTime(s.TimeStart), convertToTime(s.TimeEnd)
}
func convertToTime(timeStr string) *time.Time {
	now := time.Now()
	t, err := time.Parse("15:04", timeStr)
	if err != nil {
		return &time.Time{}
	}
	result := time.Date(now.Year(), now.Month(), now.Day(), t.Hour(), t.Minute(), 0, 0, now.Location())
	return &result
}
