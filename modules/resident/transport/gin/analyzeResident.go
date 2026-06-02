package ginResident

import (
	"github.com/gin-gonic/gin"
	"main.go/common"
	"main.go/config"
	"main.go/modules/resident/biz"
	"main.go/modules/resident/storage"
	"net/http"
)

func AnalyzeResident(cfg *config.Configuration) func(*gin.Context) {
	return func(c *gin.Context) {
		store := storage.NewResident()
		business := biz.NewResidentCommon(store, cfg)
		
		res, err := business.AnalyzeCensusQuery()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": common.ErrInternal(err)})
			return
		}
		
		c.JSON(http.StatusOK, gin.H{"data": res})
	}
}
