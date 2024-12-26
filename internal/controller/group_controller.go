package controller

import "github.com/gin-gonic/gin"

type GroupController interface {
	Create(ctx *gin.Context)
	GetGroupById(ctx *gin.Context)
	GetGroups(ctx *gin.Context)
	UpdateGroupById(ctx *gin.Context)
	DeleteGroupById(ctx *gin.Context)
}
