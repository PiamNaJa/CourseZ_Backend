package utils

type Error struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}

func NewError(status int, message string) *Error {
	return &Error{
		Status:  status,
		Message: message,
	}
}

func NotFound(message string) *Error {
	return &Error{
		Status:  404,
		Message: message,
	}
}

func Unexpected(message string) *Error {
	return &Error{
		Status:  500,
		Message: message,
	}
}

func BadRequest(message string) *Error {
	return &Error{
		Status:  400,
		Message: message,
	}
}

func Unauthorized(message string) *Error {
	return &Error{
		Status:  401,
		Message: message,
	}
}