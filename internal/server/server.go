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
	smartBinCntrl := di.SmartBinDI(db, validator)
	configCntrl := di.ConfigDI(db, validator)
	groupCntrl := di.GroupDI(db, validator)
	historyCntrl := di.HistoryDI(db)
	notificationCntrl := di.NotificationDI(db)

	router.UserRouter(gin, db, userCntrl)
	router.SmartBinRouter(gin, db, smartBinCntrl)
	router.ConfigRouter(gin, db, configCntrl)
	router.GroupRouter(gin, db, groupCntrl)
	router.HistoryRouter(gin, db, historyCntrl)
	router.NotificationRouter(gin, db, notificationCntrl)

	err := gin.Run("localhost:8080")
	helper.Err(err)
}
