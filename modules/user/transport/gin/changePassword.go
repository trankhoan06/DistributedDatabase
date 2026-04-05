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

func ChangPassword(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var data model.UpdatePass
		if err := ctx.ShouldBindJSON(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": common.ErrColumn(err)})
			return
		}
		data.UserId = ctx.MustGet(common.Current_user).(common.Requester).GetUserId()
		store := storage.NewSql(db)
		hash := common.NewSha256Hash()
		business := biz.NewRegisterCommon(store, hash)
		err := business.NewChangePasswork(ctx.Request.Context(), data)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.(*common.AppError)})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"data": true})
	}
}
