package common

import "esim/modules/user/model"

type Payload struct {
	UId  int             `json:"user_id"`
	Role *model.RoleUser `json:"role"`
}

//
//func newPayload(uId int, role *model.RoleUser) *Payload {
//	return &Payload{
//		UId:  uId,
//		Role: role,
//	}
//}

func (p *Payload) GetUser() int {
	return p.UId
}
func (p *Payload) GetRole() *model.RoleUser {
	return p.Role
}

type Requester interface {
	GetUserId() int
	GetRole() *model.RoleUser
	GetEmail() string
}
type UserIdCommon struct {
	UserId int `json:"user_id"`
}
type IdCommon struct {
	UserId int `json:"-"`
	Id     int `json:"id"`
}
