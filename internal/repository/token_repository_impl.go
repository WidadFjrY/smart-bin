package repository

import (
	"context"
	"smart-trash-bin/domain/model"
	"smart-trash-bin/pkg/exception"
	"smart-trash-bin/pkg/helper"
	"time"

	"gorm.io/gorm"
)

type TokenRepositoryImpl struct{}

func NewTokenRepository() TokenRepository {
	return &TokenRepositoryImpl{}
}

func (repo *TokenRepositoryImpl) Create(ctx context.Context, tx *gorm.DB, token model.TokenAuth) {
	result := tx.WithContext(ctx).Create(&token)
	helper.Err(result.Error)
}

func (repo *TokenRepositoryImpl) IsValid(ctx context.Context, tx *gorm.DB, token string) bool {
	var tokenModel model.TokenAuth

	result := tx.WithContext(ctx).Where("token = ?", token).Select("is_valid").First(&tokenModel).Error
	if result != nil {
		return false
	}
	return tokenModel.IsValid
}

func (repo *TokenRepositoryImpl) Update(ctx context.Context, tx *gorm.DB, token string) {
	result := tx.WithContext(ctx).Table("token_auths").Where("token = ?", token).Updates(
		map[string]interface{}{
			"is_valid":   false,
			"updated_at": time.Now(),
		})

	if result.RowsAffected == 0 {
		panic(exception.NewBadRequestError("invalid token"))
	}
	helper.Err(result.Error)
}
