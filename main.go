package main

import (
	"github.com/gin-gonic/gin"
	"main.go/config"
	ginResident "main.go/modules/resident/transport/gin"
)

func main() {
	cfg := config.GetConfig()
	r := gin.Default()

	resident := r.Group("/resident")
	{
		resident.GET("/total", ginResident.TotalResidents(cfg))
		resident.GET("/total_sequential", ginResident.TotalResidentsSequential(cfg))
		resident.GET("/analyze", ginResident.AnalyzeResident(cfg))
	}
	r.Run(":3000")

}
