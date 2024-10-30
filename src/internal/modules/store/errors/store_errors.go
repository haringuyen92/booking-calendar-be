package store_errors

import "booking-calendar-server-backend/pkg/errors"

const (
	StoreUnknownCode = iota + errors.BaseStoreErrorCode
	StoreNotFoundCode
	SettingTimeNotFoundCode
)

var (
	StoreUnknownError        = errors.NewError(StoreUnknownCode, "Store Unknown error")
	StoreNotFoundError       = errors.NewError(StoreNotFoundCode, "Store Not Found")
	SettingTimeNotFoundError = errors.NewError(StoreUnknownCode, "SettingTime Not Found")
)
