package gin

import (
	"esim/common"
	"esim/modules/user/biz"
	"esim/modules/user/model"
	"esim/modules/user/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func UpdateUser(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var l model.UpdateUser
		if err := c.ShouldBindJSON(&l); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": common.ErrColumn(err),
			})
			return
		}
		l.Email = c.MustGet(common.Current_user).(common.Requester).GetEmail()
		store := storage.NewSql(db)
		business := biz.NewUserCommon(store)
		err := business.NewUpdateUser(c.Request.Context(), &l)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.(*common.AppError),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": "success"})
	}
}
