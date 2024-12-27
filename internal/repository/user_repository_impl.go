package repository

import (
	"context"
	"errors"
	"fmt"
	"smart-trash-bin/domain/model"
	"smart-trash-bin/pkg/exception"
	"smart-trash-bin/pkg/helper"
	"strings"
	"time"

	"gorm.io/gorm"
)

type UserRepostioryImpl struct{}

func NewUserRepository() UserRepostiory {
	return &UserRepostioryImpl{}
}

func (Repo *UserRepostioryImpl) CreateUser(ctx context.Context, tx *gorm.DB, user model.User) model.User {
	result := tx.WithContext(ctx).Create(&user)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "Error 1062") {
			panic(exception.NewBadRequestError("email already in use"))
		} else {
			panic(result.Error.Error())
		}
	}

	user.CreatedAt = time.Now()
	return user
}

func (Repo *UserRepostioryImpl) GetUserByEmail(ctx context.Context, tx *gorm.DB, email string) model.User {
	var user model.User

	result := tx.WithContext(ctx).Where("email = ?", email).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		panic(exception.NewUnauthorized("email or password wrong!"))
	}

	helper.Err(result.Error)
	return user
}

func (Repo *UserRepostioryImpl) GetUserById(ctx context.Context, tx *gorm.DB, userId string) model.User {
	var user model.User

	result := tx.WithContext(ctx).Preload("SmartBin").Preload("Group.SmartBin").Preload("Notification").Where("id = ?", userId).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		panic(exception.NewNotFoundError(fmt.Sprintf("user with id %s not found", userId)))
	}

	helper.Err(result.Error)
	return user
}

func (Repo *UserRepostioryImpl) IsUserExsit(ctx context.Context, tx *gorm.DB, userId string) bool {
	var user model.User

	result := tx.WithContext(ctx).Where("id = ?", userId).First(&user)
	return errors.Is(result.Error, gorm.ErrRecordNotFound) // true if user not found
}

func (Repo *UserRepostioryImpl) UpdateUserById(ctx context.Context, tx *gorm.DB, user model.User) model.User {
	err := tx.WithContext(ctx).Table("users").Where("id = ?", user.ID).Updates(
		map[string]interface{}{
			"name":       user.Name,
			"updated_at": time.Now(),
		},
	).Error
	helper.Err(err)

	user.UpdatedAt = time.Now()
	return user
}

func (Repo *UserRepostioryImpl) UpdatePasswordById(ctx context.Context, tx *gorm.DB, user model.User) model.User {
	err := tx.WithContext(ctx).Table("users").Where("id = ?", user.ID).Updates(
		map[string]interface{}{
			"password":   user.Password,
			"updated_at": time.Now(),
		},
	).Error
	helper.Err(err)

	user.UpdatedAt = time.Now()
	return user
}
