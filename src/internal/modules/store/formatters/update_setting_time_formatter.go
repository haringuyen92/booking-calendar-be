package store_formatters

import (
	common_dto "booking-calendar-server-backend/internal/core/common_dto/store"
	store_dto "booking-calendar-server-backend/internal/modules/store/dto"
	store_requests2 "booking-calendar-server-backend/internal/modules/store/requests/settings"
)

func UpdateSettingTimeFormatter(req *store_requests2.UpdateSettingTimeRequest) *store_dto.UpdateSettingTimeDto {
	return &store_dto.UpdateSettingTimeDto{
		SettingTimeDto: common_dto.SettingTimeDto{
			StoreID:             req.StoreID,
			IsOpen:              req.IsOpen,
			IsApplyDailySetting: req.IsApplyDailySetting,
			DailySetting: common_dto.SettingTimeSlot{
				IsOpenAllDay: req.DailySetting.IsOpenAllDay,
				IsOffDay:     req.DailySetting.IsOffDay,
				SlotTime:     formatSlotTime(req.DailySetting.SlotTime),
			},
			MondaySetting: common_dto.SettingTimeSlot{
				IsOpenAllDay: req.MondaySetting.IsOpenAllDay,
				IsOffDay:     req.MondaySetting.IsOffDay,
				SlotTime:     formatSlotTime(req.MondaySetting.SlotTime),
			},
			TuesdaySetting: common_dto.SettingTimeSlot{
				IsOpenAllDay: req.TuesdaySetting.IsOpenAllDay,
				IsOffDay:     req.TuesdaySetting.IsOffDay,
				SlotTime:     formatSlotTime(req.TuesdaySetting.SlotTime),
			},
			WednesdaySetting: common_dto.SettingTimeSlot{
				IsOpenAllDay: req.WednesdaySetting.IsOpenAllDay,
				IsOffDay:     req.WednesdaySetting.IsOffDay,
				SlotTime:     formatSlotTime(req.WednesdaySetting.SlotTime),
			},
			ThursdaySetting: common_dto.SettingTimeSlot{
				IsOpenAllDay: req.ThursdaySetting.IsOpenAllDay,
				IsOffDay:     req.ThursdaySetting.IsOffDay,
				SlotTime:     formatSlotTime(req.ThursdaySetting.SlotTime),
			},
			FridaySetting: common_dto.SettingTimeSlot{
				IsOpenAllDay: req.FridaySetting.IsOpenAllDay,
				IsOffDay:     req.FridaySetting.IsOffDay,
				SlotTime:     formatSlotTime(req.FridaySetting.SlotTime),
			},
			SaturdaySetting: common_dto.SettingTimeSlot{
				IsOpenAllDay: req.SaturdaySetting.IsOpenAllDay,
				IsOffDay:     req.SaturdaySetting.IsOffDay,
				SlotTime:     formatSlotTime(req.SaturdaySetting.SlotTime),
			},
			SundaySetting: common_dto.SettingTimeSlot{
				IsOpenAllDay: req.SundaySetting.IsOpenAllDay,
				IsOffDay:     req.SundaySetting.IsOffDay,
				SlotTime:     formatSlotTime(req.SundaySetting.SlotTime),
			},
		},
	}
}

func formatSlotTime(slotTime []store_requests2.SettingTimeSlotItem) []common_dto.SettingTimeSlotItem {
	var result []common_dto.SettingTimeSlotItem
	for _, slot := range slotTime {
		timeStart, timeEnd := slot.ConvertToTime()
		result = append(result, common_dto.SettingTimeSlotItem{
			TimeStart: timeStart,
			TimeEnd:   timeEnd,
		})
	}
	return result
}
