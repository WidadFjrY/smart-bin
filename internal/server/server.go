package server

import (
	"smart-trash-bin/internal/app"
	"smart-trash-bin/internal/di"
	"smart-trash-bin/internal/middleware"
	"smart-trash-bin/pkg/helper"
	"smart-trash-bin/router"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Run() {
	logFilePaths := []string{"pkg/log/gorm.log", "pkg/log/gin.log"}
	helper.RefreshLog(logFilePaths)

	db := app.NewDB()
	validator := validator.New()

	gin := gin.Default()
	gin.Use(helper.NewFileGinLog(), middleware.ErrorHandling())

	userCntrl := di.UserDI(db, validator)
	smartBinCntrl := di.SmartBinDi(db, validator)

	router.UserRouter(gin, db, userCntrl)
	router.SmartBinRouter(gin, db, smartBinCntrl)

	err := gin.Run("localhost:8080")
	helper.Err(err)
}
