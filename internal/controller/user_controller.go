package controller

import "github.com/gin-gonic/gin"

type UserController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	GetUserById(ctx *gin.Context)
	UpdateUserById(ctx *gin.Context)
	UpdatePasswordById(ctx *gin.Context)
	LogoutUser(ctx *gin.Context)
}
