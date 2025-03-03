package dto

import (
	"github.com/pkg/errors"
)

var (
	// common errors
	ErrSuccess             = newError(0, "ok")
	ErrBadRequest          = newError(400, "Bad Request")
	ErrUnauthorized        = newError(401, "Unauthorized")
	ErrNotFound            = newError(404, "Not Found")
	ErrInternalServerError = newError(500, "Internal Server Error")
	ErrRequestParamsError  = newError(1000, "Request params error")

	// more biz errors
	ErrEmailAlreadyUse = newError(1001, "The email is already in use.")
	ErrRoleNotExit     = newError(1002, "The role is not support.")
	ErrPassword        = newError(1003, "The password is error.")
	ErrUsernameExists  = newError(1004, "The username is already exists.")
)

func CreateRequestParamsError(err error) error {
	return newError(1000, err.Error())
}

type Error struct {
	Code    int
	Message string
}

func newError(code int, msg string) error {
	err := errors.New(msg)
	errorCodeMap[err] = code
	return err
}
func (e Error) Error() string {
	return e.Message
}

var errorCodeMap = map[error]int{}
