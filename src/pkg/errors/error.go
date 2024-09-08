package errors

type Error struct {
	Code    ErrorCode
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) Equal(err error) bool {
	if err == nil {
		return false
	}

	if e.Code != err.(*Error).Code {
		return false
	}

	return e.Message == err.(*Error).Message
}
