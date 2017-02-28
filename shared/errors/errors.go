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
	// general
	ErrInvalidUsernamePassword  = NewError(1001, "Invalid username or password")
	ErrInvalidRequestData       = NewError(1002, "Invalid request data")
	ErrInternal                 = NewError(1003, "Internal server error")
	ErrInvalidRequestParameters = NewError(1004, "Missing or invalid request parameters")
	// user
	ErrUserAlreadyExists  = NewError(1005, "User already exists")
	ErrUserNotFound       = NewError(1006, "User not found")
	ErrCannotRegisterUser = NewError(1007, "Cannot register user")
	ErrInvalidUserModel   = NewError(1008, "Invalid user model")
	// category
	ErrCategoryNotFound     = NewError(1009, "Category not found")
	ErrCannotUpdateCategory = NewError(1010, "Cannot update category")
	ErrCannotRemoveCategory = NewError(1011, "Cannot remove category")
	ErrCannotGetCategories  = NewError(1012, "Cannot load categories")
	ErrInvalidCategoryModel = NewError(1013, "Invalid category model")
	// expense
	ErrExpenseNotFound     = NewError(1014, "Expense not found")
	ErrCannotUpdateExpense = NewError(1015, "Cannot update expense")
	ErrCannotRemoveExpense = NewError(1016, "Cannot remove expense")
	ErrCannotGetExpenses   = NewError(1017, "Cannot load expenses")
	ErrInvalidExpenseModel = NewError(1018, "Invalid expense model")
)
