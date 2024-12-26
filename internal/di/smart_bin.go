package di

import (
	"smart-trash-bin/internal/controller"
	"smart-trash-bin/internal/repository"
	"smart-trash-bin/internal/service"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func SmartBinDI(db *gorm.DB, validator *validator.Validate) controller.SmartBinController {
	binRepo := repository.NewSmartBinRepository()
	binServ := service.NewSmartBinService(db, validator, binRepo)

	groupRepo := repository.NewGroupRepository()
	groupServ := service.NewGroupService(db, validator, groupRepo)
	cntrl := controller.NewSmartBinController(binServ, groupServ)

	return cntrl
}
