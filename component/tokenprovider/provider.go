package tokenprovider

import (
	"esim/modules/user/model"
)

type TokenProvider interface {
	GenerateAccessToken(payload Payload) (string, error)
	GenerateRefreshToken(payload Payload) (string, error)
	//VerifyRefreshToken(ctx context.Context, payload Payload, refreshToken string, db *gorm.DB) (Token, error)
	Validate(token string) (Payload, error)
	GetSecret() string
}
type Token interface {
	GetAccessToken() string
	GetRefreshToken() string
}
type Payload interface {
	GetRole() *model.RoleUser
	GetUser() int
}
