package router

import (
	"smart-trash-bin/internal/controller"
	"smart-trash-bin/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func NotificationRouter(router *gin.Engine, db *gorm.DB, cntrl controller.NotificationController) {
	router.Use(middleware.Auth(db))
	router.Group("/")
	{
		router.PUT("/api/notification/:notif_id", cntrl.UpdateNotificationById)
		router.DELETE("/api/notification/:notif_id", cntrl.DeleteNotificaionById)
	}
}
