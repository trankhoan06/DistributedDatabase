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
func ErrInvalid(err error) *AppError {
	return NewCustomErr(err, "Invalid request", "ERRVALID")
}
func ErrEmailOfPass(err error) *AppError {
	return NewCustomErr(err, "email of pass invalid", "ERRVALID")
}
func ErrPass(err error) *AppError {
	return NewCustomErr(err, "pass invalid", "ERRVALID")
}
func ErrItem(err error) *AppError {
	return NewCustomErr(err, "item not found", "ERRITEM")
}
func ErrCart(err error) *AppError {
	return NewCustomErr(err, "cart haven't this item of this itetm has been deleted", "ERRITEM_CART")
}
func ErrUnauthorized(err error) *AppError {
	return NewFullErrorResponse(http.StatusUnauthorized, err, "Unauthorized", err.Error(), "ErrUnauthorized")
}
func ErrEmailExist(err error) *AppError {
	return NewCustomErr(err, "user already exists", "ERR_EMAIL")
}
func ErrEmailRequest(err error) *AppError {
	return NewCustomErr(err, "email is request", "ERR_EMAIL")
}
func ErrAccountRequest(err error) *AppError {
	return NewCustomErr(err, "account is request", "ERR_ACCOUNT")
}
func ErrAccountExist(err error) *AppError {
	return NewCustomErr(err, "account already exists", "ERR_ACCOUNT")
}
func ErrPermission(err error) *AppError {
	return NewCustomErr(err, "you don't permission", "ERR_PERMISSION")
}
func ErrExist(err error) *AppError {
	return NewCustomErr(err, "it already exists", "ERR_EXISTS")
}
func ErrNotExist(err error) *AppError {
	return NewCustomErr(err, "it don't exists", "ERR_NOTEXISTS")
}
func ErrSyntax(err error) *AppError {
	return NewCustomErr(err, "your syntax wrong", "ERR_SYMTAX")
}
func ErrColumn(err error) *AppError {
	return NewCustomErr(err, "Name of the column wrong or missing column", "ERR_ROW")
}
func ErrCreate(err error) *AppError {
	return NewCustomErr(err, "error when create", "ERR_CREATE")
}
func ErrDelete(err error) *AppError {
	return NewCustomErr(err, "error when delete", "ERR_DELETE")
}
func ErrGet(err error) *AppError {
	return NewCustomErr(err, "error when get", "ERR_GET")
}
func ErrUpdate(err error) *AppError {
	return NewCustomErr(err, "error when update", "ERR_UPDATE")
}
func ErrToken(err error) *AppError {
	return NewCustomErr(err, "Token wrong", "ERR_TOKEN")
}
func ErrRequest(err error) *AppError {
	return NewCustomErr(err, "Error when request", "ERR_REQUEST")
}
func ErrStatus(err error) *AppError {
	return NewCustomErr(err, "error when update status", "ERR_TOKEN")
}
func ErrFound(err error) *AppError {
	return NewCustomErr(err, "Error when find", "ERR_FOUND")
}

func ErrExpire(err error) *AppError {
	return NewCustomErr(err, "Error expire", "ERR_EXPIRE")
}

func ErrFile(err error) *AppError {
	return NewCustomErr(err, "Error file", "ERR_FILE")
}
func ErrUnmarshal(err error) *AppError {
	return NewCustomErr(err, "Error unmarshal", "ERR_UNMARSHAL")
}

func Errmarshal(err error) *AppError {
	return NewCustomErr(err, "Error marshal", "ERR_MARSHAL")
}
func ErrDataRequest(err error) *AppError {
	return NewCustomErr(err, "data request", "ERR_DATAREQUEST")
}
func ErrAccount(err error) *AppError {
	return NewCustomErr(err, "Error Account", "ERR_ACCOUNT")
}
func ErrEmail(err error) *AppError {
	return NewCustomErr(err, "Error Email", "ERR_EMAIL")
}
func ErrLogin(err error) *AppError {
	return NewCustomErr(err, "Error Login", "ERR_LOGIN")
}
func ErrRefreshToken(err error) *AppError {
	return NewCustomErr(err, "Error RefreshToken", "ERR_REFRESHTOKEN")
}
func ErrVerifyCode(err error) *AppError {
	return NewCustomErr(err, "Error verify code", "ERR_VERIFYCODE")
}
func ErrResendEmail(err error) *AppError {
	return NewCustomErr(err, "Error resend email", "ERR_RESENDCODE")
}
