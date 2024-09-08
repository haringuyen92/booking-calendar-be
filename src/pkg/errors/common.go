package errors

import "fmt"

type ErrorCode uint

const (
	BaseErrorCodeUnknown ErrorCode = iota*1000 + 1000
	BaseUserErrorCode
	BaseStoreErrorCode
	BaseCourseErrorCode
	BaseStaffErrorCode
	BaseBookingErrorCode
)

func (e ErrorCode) String() string {
	switch {
	case e >= BaseUserErrorCode && e < BaseStaffErrorCode:
		return "User"
	case e >= BaseStoreErrorCode && e < BaseCourseErrorCode:
		return "Store"
	case e >= BaseCourseErrorCode && e < BaseStaffErrorCode:
		return "Store"
	case e >= BaseStaffErrorCode && e < BaseBookingErrorCode:
		return "Course"
	case e >= BaseBookingErrorCode:
		return "Staff"
	default:
		return "Unknown"
	}
}

func messagePrefixByErrorCode(code ErrorCode, rawMsg string) (newMsg string) {
	newMsg = fmt.Sprintf("[Module %s]: %s", code.String(), rawMsg)
	return
}

func NewError(code ErrorCode, err error) *Error {
	message := messagePrefixByErrorCode(code, err.Error())
	return &Error{
		Code:    code,
		Message: message,
	}
}
