package router

import (
	"smart-trash-bin/internal/controller"
	"smart-trash-bin/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SmartBinRouter(router *gin.Engine, db *gorm.DB, cntrl controller.SmartBinController) {
	router.POST("/api/bin/classification", cntrl.ClassifyImage)
	router.POST("/api/bin/status", cntrl.UpdateSmartBinValue)
	auth := router.Group("/")
	auth.Use(middleware.Auth(db))
	{
		auth.POST("/api/bin/add", cntrl.AddSmartBin)
		auth.PUT("/api/bin/id/:bin_id", cntrl.UpdateSmartBinById)
		auth.GET("/api/bin/id/:bin_id", cntrl.GetSmartBinById)
		auth.DELETE("/api/bin/id/:bin_id", cntrl.DeleteSmartBinById)
		auth.GET("/api/bin/page/:page", cntrl.GetSmartBins)
		auth.POST("/api/bin/lock/:bin_id", cntrl.LockSmartBin)
		auth.POST("/api/bin/unlock/:bin_id", cntrl.UnlockSmartBin)
	}
}
