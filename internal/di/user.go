package di

import (
	"smart-trash-bin/internal/controller"
	"smart-trash-bin/internal/repository"
	"smart-trash-bin/internal/service"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func UserDI(db *gorm.DB, validator *validator.Validate) controller.UserController {
	repo := repository.NewUserRepository()
	serv := service.NewUserService(db, validator, repo)
	cntrl := controller.NewController(serv)

	return cntrl
}
