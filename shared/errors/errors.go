package errors

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	BaseURL = "https://et.eu/api/documentation/errors?id="
)

type Error struct {
	Code             int    `json:"code,omitempty"`
	Message          string `json:"message,omitempty"`
	DocumentationURL string `json:"documentationUrl,omitempty"`
}

func NewError(code int, message string) *Error {
	return &Error{
		Code:             code,
		Message:          message,
		DocumentationURL: BaseURL + strconv.Itoa(code),
	}
}

func (e Error) Error() string {
	s := fmt.Sprintf("%s (code: #%d)", e.Message, e.Code)

	return s
}

func (e Error) IsSame(err error) bool {
	if strings.Contains(err.Error(), e.Error()) {
		return true
	}
	return false
}

var (
	ErrUserAlreadyExists        = NewError(1000, "User already exists")
	ErrInvalidUsernamePassword  = NewError(1001, "Invalid username or password")
	ErrInvalidRequestData       = NewError(1002, "Invalid request data")
	ErrInternal                 = NewError(1003, "Internal server error")
	ErrCannotRegisterUser       = NewError(1004, "Cannot register user")
	ErrInvalidRequestParameters = NewError(1005, "Missing or invalid request parameters")
	ErrUserNotFound             = NewError(1006, "User not found")
	ErrCannotUpdateExpense      = NewError(1007, "Cannot update expense")
	ErrCannotUpdateCategory     = NewError(1007, "Cannot update category")
)
