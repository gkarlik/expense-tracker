package errors

import (
	"strings"
)

type Error string

func (e Error) Error() string { return string(e) }

func (e Error) IsSame(err error) bool {
	if e == err || strings.Contains(err.Error(), e.Error()) {
		return true
	}
	return false
}

const (
	ErrUserExists              = Error("User already exists")
	ErrUserServiceConnection   = Error("Cannot connect to UserService")
	ErrInvalidUsernamePassword = Error("Invalid username or password")
)
