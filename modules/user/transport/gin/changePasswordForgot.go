package gin

import (
	"esim/common"
	StorageEmail "esim/modules/email/storage"
	"esim/modules/user/biz"
	"esim/modules/user/model"
	"esim/modules/user/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func ChangeForgotPassword(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {
		var data model.NewPasswordForgot
		if err := c.ShouldBindJSON(&data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := storage.NewSql(db)
		storeEmail := StorageEmail.NewSqlModel(db)
		hash := common.NewSha256Hash()
		business := biz.NewVerifyCommon(store, storeEmail, hash)
		if err := business.NewChangePasswordForgot(c.Request.Context(), &data); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "success"})

	}
}
