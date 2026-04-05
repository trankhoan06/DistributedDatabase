package gin

import (
	"esim/common"
	"esim/component/tokenprovider"
	modelRefreshToken "esim/modules/refreshToken/model"
	storageRefreshToken "esim/modules/refreshToken/storage"
	"esim/modules/user/biz"
	"esim/modules/user/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func GetAccessToken(db *gorm.DB, provider tokenprovider.TokenProvider) func(*gin.Context) {
	return func(c *gin.Context) {
		var l modelRefreshToken.GetRefreshToken
		if err := c.ShouldBindJSON(&l); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": common.ErrColumn(err),
			})
			return
		}

		store := storage.NewSql(db)
		storeRef := storageRefreshToken.NewSqlModel(db)
		business := biz.NewRefreshTokenCommon(store, storeRef, provider)
		token, err := business.NewGetAccessToken(c, l)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.(*common.AppError),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": token})
	}
}
