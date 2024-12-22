package repository

import (
	"context"
	"smart-trash-bin/domain/model"

	"gorm.io/gorm"
)

type TokenRepository interface {
	Create(ctx context.Context, tx *gorm.DB, token model.TokenAuth)
	IsValid(ctx context.Context, tx *gorm.DB, token string) bool
	Update(ctx context.Context, tx *gorm.DB, token string)
}
