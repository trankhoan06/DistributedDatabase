package jwt

import (
	"errors"
	"esim/common"
	"esim/component/tokenprovider"
	modelUser "esim/modules/user/model"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

const (
	timeAcessToken   = 60 * 24 * 30
	TimeRefrestToken = 15
)

type JwtProvider struct {
	Secret string
	Prefix string
}
type token struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	Created      time.Time `json:"created"`
	Expiry       time.Time `json:"expiry"`
}
type ResponToken struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	//Role         int    `json:"role"`
	User modelUser.UserSimple `json:"user"`
}
type MyClaim struct {
	Payload common.Payload `json:"payload"`
	jwt.StandardClaims
}

func NewJwtProvider(secret string, prefix string) *JwtProvider {
	return &JwtProvider{Secret: secret, Prefix: prefix}
}
func (t *token) GetAccessToken() string {
	return t.AccessToken
}
func (t *token) GetRefreshToken() string {
	return t.RefreshToken
}
func (j *JwtProvider) GetSecret() string {
	return j.Secret
}
func (j *JwtProvider) GenerateAccessToken(payload tokenprovider.Payload) (string, error) {
	now := time.Now()
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, &MyClaim{
		Payload: common.Payload{
			UId:  payload.GetUser(),
			Role: payload.GetRole(),
		},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: now.Add(time.Duration(timeAcessToken) * time.Minute).Unix(),
			IssuedAt:  now.Unix(),
			Issuer:    fmt.Sprint(now.UnixNano()),
		},
	})
	MyToken, err := t.SignedString([]byte(j.Secret))
	if err != nil {
		return "", err
	}
	return MyToken, nil
}
func (j *JwtProvider) GenerateRefreshToken(payload tokenprovider.Payload) (string, error) {
	now := time.Now()
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, &MyClaim{
		Payload: common.Payload{
			UId:  payload.GetUser(),
			Role: payload.GetRole(),
		},
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: now.Add(time.Duration(0) * time.Minute).Unix(),
			IssuedAt:  now.Unix(),
			Issuer:    fmt.Sprint(now.UnixNano()),
		},
	})
	MyToken, err := t.SignedString([]byte(j.Secret))
	if err != nil {
		return "", err
	}
	return MyToken, nil
}

//	func (j *JwtProvider) VerifyRefreshToken(ctx context.Context, payload tokenprovider.Payload, refreshToken string, db *gorm.DB) (tokenprovider.Token, error) {
//		store := storageRefreshToken.NewSqlModel(db)
//		business := tokenprovider.NewRefreshTokenBiz(store)
//		now := time.Now().Add(-7 * time.Hour)
//		data, err := business.RefreshTokenInter.FindRefreshToken(ctx, refreshToken)
//		if err != nil {
//			return nil, err
//		}
//		if data == nil {
//			return nil, errors.New("refresh token is invalid")
//		}
//		if data.ExpireAt.After(now) {
//			return nil, errors.New("refresh token is expired")
//		}
//		MyToken, err := j.GenerateAccessToken(payload)
//		if err != nil {
//			return nil, err
//		}
//		return &token{
//			AccessToken:  MyToken,
//			RefreshToken: refreshToken,
//			Expiry:       time.Now().Add(time.Duration(timeAcessToken) * time.Minute),
//			Created:      now,
//		}, nil
//	}
func (j *JwtProvider) Validate(token string) (tokenprovider.Payload, error) {
	MyToken, err := jwt.ParseWithClaims(token, &MyClaim{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(j.Secret), nil

	})
	if err != nil {
		return nil, err
	}
	if MyToken == nil {
		return nil, errors.New("token is nil")
	}
	if !MyToken.Valid {
		return nil, errors.New("token is invalid")
	}
	claim, ok := MyToken.Claims.(*MyClaim)
	if !ok {
		return nil, errors.New("token is invalid")
	}
	return &claim.Payload, nil
}
