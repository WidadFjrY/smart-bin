package di

import (
	"smart-trash-bin/internal/controller"
	"smart-trash-bin/internal/repository"
	"smart-trash-bin/internal/service"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func ConfigDI(db *gorm.DB, validator *validator.Validate) controller.ConfigController {
	configRepo := repository.NewConfigRepository()
	configServ := service.NewConfigService(db, validator, configRepo)
	configCntrl := controller.NewConfigController(configServ)

	return configCntrl
}
