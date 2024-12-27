package controller

import "github.com/gin-gonic/gin"

type NotificationController interface {
	UpdateNotificationById(ctx *gin.Context)
	DeleteNotificaionById(ctx *gin.Context)
}
