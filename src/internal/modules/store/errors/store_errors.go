package store_errors

import "booking-calendar-server-backend/pkg/errors"

const (
	StoreUnknownCode = iota + errors.BaseStoreErrorCode
	StoreNotFoundCode
	SettingTimeNotFoundCode
	SettingBookingNotFoundCode
	SettingSlotNotFoundCode
)

var (
	StoreUnknownError           = errors.NewError(StoreUnknownCode, "Store Unknown error")
	StoreNotFoundError          = errors.NewError(StoreNotFoundCode, "Store Not Found")
	SettingTimeNotFoundError    = errors.NewError(StoreUnknownCode, "SettingTime Not Found")
	SettingBookingNotFoundError = errors.NewError(StoreUnknownCode, "SettingBooking Not Found")
	SettingSlotNotFoundError    = errors.NewError(StoreUnknownCode, "SettingSlot Not Found")
)
