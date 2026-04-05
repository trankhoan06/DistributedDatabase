package ginMail

import (
	"esim/common"
	"esim/modules/email/biz"
	"esim/modules/email/model"
	"esim/modules/email/storage"
	storageUser "esim/modules/user/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func VerifyForgotPassword(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.VerifyAccountCode
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": common.ErrColumn(err)})
			return
		}
		store := storage.NewSqlModel(db)
		storeUser := storageUser.NewSql(db)
		business := biz.NewSendEmailBiz(store, storeUser)
		if err := business.NewVerifyForgotPassword(c.Request.Context(), &data, 60*5); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.(*common.AppError)})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": true})
	}
}
