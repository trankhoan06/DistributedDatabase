package common

import (
	"errors"
	"net/http"
)

type AppError struct {
	StatusCode int    `json:"status_code"`
	Root       error  `json:"-"`
	Msg        string `json:"msg"`
	Log        string `json:"log"`
	Key        string `json:"key"`
}

func NewAppError(root error, msg string, log string, key string) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		Root:       root,
		Msg:        msg,
		Log:        log,
		Key:        key,
	}
}
func (e *AppError) RootErr() error {
	if err, ok := e.Root.(*AppError); ok {
		return err.RootErr()
	}
	return e.Root
}
func (e *AppError) Error() string {
	return e.RootErr().Error()
}
func NewFullErrorResponse(status int, root error, msg, log, key string) *AppError {
	return &AppError{
		status,
		root,
		msg,
		log,
		key,
	}
}
func NewAuthorize(root error, msg, log, key string) *AppError {
	return &AppError{
		http.StatusUnauthorized,
		root,
		msg,
		log,
		key,
	}
}
func ErrInternal(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err,
		"Something went wrong in the server", err.Error(), "ErrInternal")
}
func NewCustomErr(root error, msg, key string) *AppError {
	if root != nil {
		return NewAppError(root, msg, root.Error(), key)
	}
	return NewAppError(errors.New(msg), msg, msg, key)
}
func ErrDb(err error) *AppError {
	return NewFullErrorResponse(http.StatusInternalServerError, err, "Something went wrong with DB", err.Error(), "DB_ERROR")
}
func ErrServer(err error) *AppError {
	return NewCustomErr(err, "Invalid request", "ERRVALID")
}
