package di

import (
	"smart-trash-bin/internal/controller"
	"smart-trash-bin/internal/repository"
	"smart-trash-bin/internal/service"

	"gorm.io/gorm"
)

func HistoryDI(db *gorm.DB) controller.HistoryController {
	repo := repository.NewHistoryRepository()
	serv := service.NewHistoryService(db, repo)
	cntrl := controller.NewHistoryController(serv)

	return cntrl
}
