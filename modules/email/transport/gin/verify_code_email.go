package ginMail

import (
	"esim/common"
	"esim/component/tokenprovider"
	"esim/modules/email/biz"
	"esim/modules/email/model"
	"esim/modules/email/storage"
	storageRefreshToken "esim/modules/refreshToken/storage"
	storageUser "esim/modules/user/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func VerifyCodeEmail(db *gorm.DB, provider tokenprovider.TokenProvider) func(*gin.Context) {
	return func(c *gin.Context) {
		var verify model.VerifyAccountCode
		if err := c.ShouldBindJSON(&verify); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": common.ErrColumn(err)})
			return
		}
		store := storage.NewSqlModel(db)
		storeUser := storageUser.NewSql(db)
		storerefresh := storageRefreshToken.NewSqlModel(db)
		hash := common.NewSha256Hash()
		business := biz.NewLoginBiz(store, storeUser, provider, hash, storerefresh)
		token, err := business.NewVerifyEmail(c.Request.Context(), &verify)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.(*common.AppError)})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": token})
	}
}
