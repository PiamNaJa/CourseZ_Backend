package utils

import (
	"errors"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

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
	return NewError(fiber.StatusNotFound, message)
}

func Unexpected(message string) *Error {
	return NewError(fiber.StatusInternalServerError, message)
}

func BadRequest(message string) *Error {
	return NewError(fiber.StatusBadRequest, message)
}

func Unauthorized(message string) *Error {
	return NewError(fiber.StatusUnauthorized, message)
}

func HandleRecordNotFoundErr(err error) *Error {
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return NotFound(err.Error())
	}
	return Unexpected(err.Error())
}
