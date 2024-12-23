package di

import (
	"smart-trash-bin/internal/controller"
	"smart-trash-bin/internal/repository"
	"smart-trash-bin/internal/service"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func SmartBinDI(db *gorm.DB, validator *validator.Validate) controller.SmartBinController {
	repo := repository.NewSmartBinRepository()
	serv := service.NewSmartBinService(db, validator, repo)
	cntrl := controller.NewSmartBinController(serv)

	return cntrl
}
