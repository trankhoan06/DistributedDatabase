package gin

import (
	"esim/common"
	"esim/component/tokenprovider"
	"esim/config"
	storageRefreshToken "esim/modules/refreshToken/storage"
	"esim/modules/user/biz"
	"esim/modules/user/model"
	"esim/modules/user/storage"
	"esim/worker"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func Login(db *gorm.DB, cfg *config.Configuration, provider tokenprovider.TokenProvider, distributor worker.TaskDistributor) func(*gin.Context) {
	return func(c *gin.Context) {
		var l model.Login
		if err := c.ShouldBindJSON(&l); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": common.ErrColumn(err),
			})
			return
		}
		store := storage.NewSql(db)
		hash := common.NewSha256Hash()
		storeRef := storageRefreshToken.NewSqlModel(db)
		business := biz.NewLoginBiz(store, storeRef, hash, cfg, provider, distributor)
		token, err := business.NewLogin(c, &l, db)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.(*common.AppError),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": token})
	}
}
