package router

import (
	"smart-trash-bin/internal/controller"
	"smart-trash-bin/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func HistoryRouter(router *gin.Engine, db *gorm.DB, cntrl controller.HistoryController) {
	auth := router.Group("/")
	auth.Use(middleware.Auth(db))
	{
		auth.GET("/api/history/bin_id/:bin_id", cntrl.GetHistoriesByBinId)
	}
}
