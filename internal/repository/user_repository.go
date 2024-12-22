package repository

import (
	"context"
	"smart-trash-bin/domain/model"

	"gorm.io/gorm"
)

type UserRepostiory interface {
	CreateUser(ctx context.Context, tx *gorm.DB, user model.User) model.User
	GetUserByEmail(ctx context.Context, tx *gorm.DB, email string) model.User
	GetUserById(ctx context.Context, tx *gorm.DB, userId string) model.User
	IsUserExsit(ctx context.Context, tx *gorm.DB, userId string) bool
	UpdateUserById(ctx context.Context, tx *gorm.DB, user model.User) model.User
	UpdatePasswordById(ctx context.Context, tx *gorm.DB, user model.User) model.User
}
