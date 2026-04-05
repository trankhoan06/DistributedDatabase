package gin

import (
	"esim/modules/user/biz"
	"esim/modules/user/model"
	"esim/modules/user/storage"
	"esim/worker"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func ForgotPassword(db *gorm.DB, taskDistributor worker.TaskDistributor) func(c *gin.Context) {
	return func(c *gin.Context) {
		var forgot model.ForgotPassword
		if err := c.ShouldBindJSON(&forgot); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		store := storage.NewSql(db)
		business := biz.NewForgotCommon(store, taskDistributor)
		err := business.NewForgotPassword(c.Request.Context(), &forgot)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"data": "success"})
	}
}
