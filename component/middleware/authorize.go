package middleware

import (
	"context"
	"errors"
	"esim/common"
	"esim/modules/user/model"
	"github.com/gin-gonic/gin"
	"log"
	"strings"
)

type Authorize interface {
	FindUser(ctx context.Context, cond map[string]interface{}) (*model.User, error)
}

func ExactlyToken(s string) (string, error) {
	str := strings.Split(s, " ")
	// "Auth": Bear {token}
	if str[0] != "Bearer" || len(str) < 2 || strings.TrimSpace(str[1]) == "" {
		return "", errors.New("token has been fault")
	}
	return str[1], nil
}
func (j *ModelMiddleware) RequestAuthorize() func(c *gin.Context) {
	return func(c *gin.Context) {
		s, err := ExactlyToken(c.GetHeader("Authorization"))
		if err != nil {
			log.Println("Exactly_token err:", err)
			panic(common.ErrUnauthorized(err))
		}
		payLoad, err := j.token.Validate(s)
		if err != nil {
			log.Println("token err", err)
			panic(common.ErrUnauthorized(err))
		}
		user, err1 := j.authen.FindUser(c.Request.Context(), map[string]interface{}{"id": payLoad.GetUser()})
		if err1 != nil {
			log.Println("token error", err1, payLoad)
			panic(common.ErrUnauthorized(err))
		}
		c.Set(common.Current_user, user)
		c.Next()

	}
}
