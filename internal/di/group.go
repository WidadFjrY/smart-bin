package di

import (
	"smart-trash-bin/internal/controller"
	"smart-trash-bin/internal/repository"
	"smart-trash-bin/internal/service"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

func GroupDI(db *gorm.DB, validator *validator.Validate) controller.GroupController {
	repo := repository.NewGroupRepository()
	serv := service.NewGroupService(db, validator, repo)
	cntrl := controller.NewGroupController(serv)

	return cntrl
}
