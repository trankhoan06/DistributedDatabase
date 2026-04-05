package common

import "time"

type RefreshToken struct {
	Id       int       `json:"id" gorm:"column:id"`
	UserId   int       `json:"user_id" gorm:"column:user_id"`
	Token    string    `json:"token" gorm:"column:token"`
	CreateAt time.Time `json:"create_at" gorm:"column:create_at"`
	ExpireAt time.Time `json:"expire_at" gorm:"column:expire_at"`
}
type CreateRefreshToken struct {
	UserId   int       `json:"-" gorm:"column:user_id"`
	Token    string    `json:"token" gorm:"column:token"`
	ExpireAt time.Time `json:"expire_at" gorm:"column:expire_at"`
}
type GetRefreshToken struct {
	Token string `json:"refresh_token" gorm:"column:token"`
}

func (RefreshToken) TableName() string       { return "refresh_token" }
func (CreateRefreshToken) TableName() string { return "refresh_token" }
