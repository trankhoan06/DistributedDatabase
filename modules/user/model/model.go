package model

import "time"

type RoleUser int

const (
	RoleUserUser RoleUser = iota
	RoleUserAdmin
)

type User struct {
	Id        int       `json:"id" gorm:"column:id"`
	Account   string    `json:"account" gorm:"column:account"`
	Password  string    `json:"password" gorm:"column:password"`
	Salt      string    `json:"-" gorm:"column:salt"`
	Email     string    `json:"email" gorm:"column:email"`
	Role      *RoleUser `json:"role" gorm:"column:role"`
	FirstName string    `json:"first_name" gorm:"column:first_name"`
	LastName  string    `json:"last_name" gorm:"column:last_name"`
	CreateAt  time.Time `json:"create_at" gorm:"column:create_at"`
	UpdateAt  time.Time `json:"update_at" gorm:"column:update_at"`
}
type UserSimple struct {
	Id    int       `json:"id" gorm:"column:id"`
	Email string    `json:"email" gorm:"column:email"`
	Role  *RoleUser `json:"role" gorm:"column:role"`
}
type UpdateUser struct {
	Email     string  `json:"-" gorm:"column:email"`
	FirstName *string `json:"first_name" gorm:"column:first_name"`
	LastName  *string `json:"last_name" gorm:"column:last_name"`
}
type CreateUser struct {
	Account   string `json:"account" gorm:"column:account"`
	Password  string `json:"password" gorm:"column:password"`
	Salt      string `json:"-" gorm:"column:salt"`
	Token     string `json:"-" gorm:"column:token"`
	Email     string `json:"email" gorm:"column:email"`
	FirstName string `json:"first_name" gorm:"column:first_name"`
	LastName  string `json:"last_name" gorm:"column:last_name"`
}

type Login struct {
	Account  string `json:"account" gorm:"column:account"`
	Password string `json:"password" gorm:"column:password"`
	DeviceId string `json:"device_id" gorm:"column:device_id"`
}
type UpdatePass struct {
	OldPassword string `json:"old_password" gorm:"column:old_password"`
	NewPassword string `json:"new_password" gorm:"column:new_password"`
	UserId      int    `json:"-" gorm:"column:user_id"`
}

type NewPasswordForgot struct {
	NewPassword string `json:"new_password" gorm:"column:new_password"`
	Email       string `json:"email" gorm:"column:email"`
}

type ForgotPassword struct {
	Email string `json:"email" gorm:"column:email"`
}

func (u *User) GetUserId() int {
	return u.Id
}
func (u *User) GetRole() *RoleUser {
	return u.Role
}
func (u *User) GetEmail() string {
	return u.Email
}
func (User) TableName() string       { return "user" }
func (CreateUser) TableName() string { return "user" }
func (UpdateUser) TableName() string { return "user" }
