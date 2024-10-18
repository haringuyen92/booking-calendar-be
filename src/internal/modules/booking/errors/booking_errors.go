package booking_errors

import "booking-calendar-server-backend/pkg/errors"

const (
	BookingUnknownCode = iota + errors.BaseBookingErrorCode
	BookingNotFoundCode
	BookingAlreadyExistsCode
	BookingInvalidCode
	BookingNotAuthorizedCode
	BookingNotAcceptableCode
	BookingNotAllowedCode
	BookingNotModifiedCode
	BookingNotImplementedCode
)

var (
	BookingUnknownError        = errors.NewError(BookingUnknownCode, "BookingUnknown")
	BookingNotFoundError       = errors.NewError(BookingNotFoundCode, "BookingNotFound")
	BookingAlreadyExistsError  = errors.NewError(BookingAlreadyExistsCode, "BookingAlreadyExists")
	BookingInvalidError        = errors.NewError(BookingInvalidCode, "BookingInvalid")
	BookingNotAuthorizedError  = errors.NewError(BookingNotAuthorizedCode, "BookingInvalid")
	BookingNotAcceptableError  = errors.NewError(BookingNotAcceptableCode, "BookingInvalid")
	BookingNotAllowedError     = errors.NewError(BookingNotAllowedCode, "BookingInvalid")
	BookingNotModifiedError    = errors.NewError(BookingNotModifiedCode, "BookingInvalid")
	BookingNotImplementedError = errors.NewError(BookingNotImplementedCode, "BookingInvalid")
)
