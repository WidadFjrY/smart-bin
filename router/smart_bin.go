package router

import (
	"smart-trash-bin/internal/controller"
	"smart-trash-bin/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SmartBinRouter(router *gin.Engine, db *gorm.DB, cntrl controller.SmartBinController) {
	router.POST("/api/bin/classification", cntrl.ClassifyImage)
	router.PUT("/api/bin/status", cntrl.UpdateSmartBinValue)

	auth := router.Group("/")
	auth.Use(middleware.Auth(db))
	{
		auth.POST("/api/bin/add", cntrl.AddSmartBin)
		auth.PUT("/api/bin/id/:bin_id", cntrl.UpdateSmartBinById)
		auth.GET("/api/bin/id/:bin_id", cntrl.GetSmartBinById)
		auth.DELETE("/api/bin/id/:bin_id", cntrl.DeleteSmartBinById)
		auth.GET("/api/bin/page/:page", cntrl.GetSmartBins)
		auth.PUT("/api/bin/lock/:bin_id", cntrl.LockSmartBin)
		auth.PUT("/api/bin/unlock/:bin_id", cntrl.UnlockSmartBin)
		auth.PUT("/api/group/add_bin/:group_id", cntrl.AddSmartBinToGroup)
		auth.PUT("/api/group/remove_bin/", cntrl.RemoveSmartBinFromGroup)
		auth.PUT("/api/group/lock/:group_id", cntrl.LockByGroup)
		auth.PUT("/api/group/unlock/:group_id", cntrl.UnlockByGroup)
	}
}
