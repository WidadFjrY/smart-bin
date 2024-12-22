package router

import (
	"smart-trash-bin/internal/controller"
	"smart-trash-bin/internal/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func UserRouter(router *gin.Engine, db *gorm.DB, cntrl controller.UserController) {
	router.POST("/api/user/register", cntrl.Register)
	router.POST("/api/user/login", cntrl.Login)

	auth := router.Group("/")
	auth.Use(middleware.Auth(db))
	{
		auth.GET("/api/user", cntrl.GetUserById)
		auth.PUT("/api/user", cntrl.UpdateUserById)
		auth.PUT("/api/user/password", cntrl.UpdatePasswordById)
		auth.POST("/api/user/logout", cntrl.LogoutUser)
	}
}
