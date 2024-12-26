package router

import (
	"smart-trash-bin/internal/controller"
	"smart-trash-bin/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func GroupRouter(router *gin.Engine, db *gorm.DB, cntrl controller.GroupController) {
	auth := router.Group("/")
	auth.Use(middleware.Auth(db))
	{
		auth.POST("/api/group", cntrl.Create)
		auth.GET("/api/group/id/:group_id", cntrl.GetGroupById)
		auth.GET("/api/group/:page", cntrl.GetGroups)
		auth.PUT("/api/group/id/:group_id", cntrl.UpdateGroupById)
	}
}
