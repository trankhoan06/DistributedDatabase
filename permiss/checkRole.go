package permiss

import (
	"errors"
	"esim/common"
	modelUser "esim/modules/user/model"
	"github.com/gin-gonic/gin"
	"log"
)

func CheckRole() func(c *gin.Context) {
	return func(c *gin.Context) {
		if *c.MustGet(common.Current_user).(common.Requester).GetRole() != modelUser.RoleUserAdmin {
			log.Println("Exactly_token err:", "user don't host")
			panic(common.ErrPermission(errors.New("user don't host")))
		}

	}
}
