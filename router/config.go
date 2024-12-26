package router

import (
	"smart-trash-bin/internal/controller"
	"smart-trash-bin/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func ConfigRouter(router *gin.Engine, db *gorm.DB, cntrl controller.ConfigController) {
	auth := router.Group("/")
	auth.Use(middleware.Auth(db))
	{
		auth.PUT("/api/config/bin_id/:bin_id", cntrl.UpdateConfigById)
	}
}
