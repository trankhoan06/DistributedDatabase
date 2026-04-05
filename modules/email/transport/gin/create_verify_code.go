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

func CreateVerifyCodeEmail(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var verify model.CreateVerifyAccount
		if err := c.ShouldBindJSON(&verify); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": common.ErrColumn(err)})
			return
		}
		store := storage.NewSqlModel(db)
		storeUser := storageUser.NewSql(db)
		business := biz.NewSendEmailBiz(store, storeUser)
		createVerify, err := business.NewResendCodeEmail(c.Request.Context(), verify.Email, 5*60, verify.Type)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.(*common.AppError)})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": createVerify})
	}
}
