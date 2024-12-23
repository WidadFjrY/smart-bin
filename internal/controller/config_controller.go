package controller

import "github.com/gin-gonic/gin"

type ConfigController interface {
	UpdateConfigById(ctx *gin.Context)
}
