package gin

import (
	"esim/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func GetProfile(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		u := c.MustGet(common.Current_user).(common.Requester)
		ip := c.ClientIP()
		fmt.Println("RemoteAddr:", ip)
		ip1 := c.Request.RemoteAddr
		fmt.Println("RemoteAddr:", ip1)
		c.JSON(http.StatusOK, gin.H{"data": u, "ip": ip})
	}
}
