package gin

import (
	"errors"
	"esim/common"
	"esim/modules/user/biz"
	"esim/modules/user/model"
	"esim/modules/user/storage"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func Register(db *gorm.DB) func(*gin.Context) {
	return func(ctx *gin.Context) {
		var data model.CreateUser
		if err := ctx.ShouldBindJSON(&data); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": common.ErrColumn(err)})
			return
		}
		if *ctx.MustGet(common.Current_user).(common.Requester).GetRole() != model.RoleUserAdmin {
			ctx.JSON(http.StatusBadRequest, common.ErrPermission(errors.New("You do not have permission to create an account")))
			return
		}
		store := storage.NewSql(db)
		hash := common.NewSha256Hash()
		business := biz.NewRegisterCommon(store, hash)
		err := business.NewRegister(ctx.Request.Context(), &data)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.(*common.AppError)})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"data": true})
	}
}
