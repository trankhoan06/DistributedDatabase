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
	}
	//ctx, cancel := context.WithCancel(context.Background())
	//worker.InitEmailWorker(ctx, sender, 100)
	//go func() {
	//	sig := make(chan os.Signal, 1)
	//	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	//	<-sig
	//	cancel()
	//}()
	r.Run(":3000")
}
