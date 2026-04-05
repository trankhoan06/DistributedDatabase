package model

import "time"

type Devide struct {
	Id       int       `json:"id" gorm:"column:id"`
	UserId   int       `json:"user_id" gorm:"column:user_id"`
	DeviceId string    `json:"device_id" gorm:"column:device_id"`
	CreateAt time.Time `json:"create_at" gorm:"column:create_at"`
	UpdateAt time.Time `json:"update_at" gorm:"column:update_at"`
}
type CreateDevide struct {
	UserId   int       `json:"-" gorm:"column:user_id"`
	DeviceId string    `json:"device_id" gorm:"column:device_id"`
	UpdateAt time.Time `json:"update_at" gorm:"column:update_at"`
}

func (CreateDevide) TableName() string { return "device" }
