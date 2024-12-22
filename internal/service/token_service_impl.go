package service

import (
	"context"
	"smart-trash-bin/domain/model"
	"smart-trash-bin/internal/repository"
	"smart-trash-bin/pkg/helper"

	"gorm.io/gorm"
)

type TokenServiceImpl struct {
	DB        *gorm.DB
	TokenRepo repository.TokenRepository
}

func NewTokenService(db *gorm.DB, tokenRepo repository.TokenRepository) TokenService {
	return &TokenServiceImpl{DB: db, TokenRepo: tokenRepo}
}

func (serv *TokenServiceImpl) Create(ctx context.Context, token string) {
	tokenModel := model.TokenAuth{
		Token: token,
	}

	txErr := serv.DB.Transaction(func(tx *gorm.DB) error {
		serv.TokenRepo.Create(ctx, tx, tokenModel)
		return nil
	})
	helper.Err(txErr)
}

func (serv *TokenServiceImpl) IsValid(ctx context.Context, token string) bool {
	var isValid bool
	serv.DB.Transaction(func(tx *gorm.DB) error {
		isValid = serv.TokenRepo.IsValid(ctx, tx, token)
		return nil
	})
	return isValid
}

func (serv *TokenServiceImpl) Update(ctx context.Context, token string) {
	txErr := serv.DB.Transaction(func(tx *gorm.DB) error {
		serv.TokenRepo.Update(ctx, tx, token)
		return nil
	})
	helper.Err(txErr)
}
