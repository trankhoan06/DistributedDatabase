package model

import (
	"time"
)

type TypeCode int

const (
	TypeVerifyEmail TypeCode = iota
	TypeForgotPassword
)

type VerifyAccount struct {
	Id       int       `json:"id" gorm:"id"`
	Email    string    `json:"email" gorm:"email"`
	Code     int       `json:"code" gorm:"code"`
	Type     *TypeCode `json:"type" gorm:"type"`
	Verify   bool      `json:"verify" gorm:"verify"`
	CreateAt time.Time `json:"create_at" gorm:"create_at"`
	Expire   time.Time `json:"expire" gorm:"expire"`
}
type CreateVerifyAccount struct {
	Email  string    `json:"email" gorm:"email"`
	Verify bool      `json:"verify" gorm:"verify"`
	Type   *TypeCode `json:"type" gorm:"type"`
	Code   int       `json:"-" gorm:"code"`
	Expire time.Time `json:"expire" gorm:"expire"`
}
type VerifyAccountCode struct {
	Email string `json:"email" gorm:"email"`
	Code  int    `json:"code" gorm:"code"`
}
type ResponseVerifyAccount struct {
	Token       string `json:"refresh_token" gorm:"column:token"`
	AccessToken string `json:"access_token" gorm:"column:access_token"`
	DeviceId    string `json:"device_id" gorm:"column:device_id"`
}

func (CreateVerifyAccount) TableName() string { return "send_code" }
func (VerifyAccount) TableName() string       { return "send_code" }
