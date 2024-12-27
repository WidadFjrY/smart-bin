package di

import (
	"smart-trash-bin/internal/controller"
	"smart-trash-bin/internal/repository"
	"smart-trash-bin/internal/service"

	"gorm.io/gorm"
)

func NotificationDI(db *gorm.DB) controller.NotificationController {
	repo := repository.NewNotificationRepository()
	serv := service.NewNotificationService(db, repo)
	cntrl := controller.NewNotificaitonController(serv)

	return cntrl
}
