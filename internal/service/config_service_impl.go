package service

import (
	"context"
	"math/rand"
	"smart-trash-bin/domain/model"
	"smart-trash-bin/domain/web"
	"smart-trash-bin/internal/repository"
	"smart-trash-bin/pkg/helper"
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type ConfigServiceImpl struct {
	DB         *gorm.DB
	Validator  *validator.Validate
	ConfigRepo repository.ConfigRepository
}

func NewConfigService(db *gorm.DB, validator *validator.Validate, configRepo repository.ConfigRepository) ConfigService {
	return &ConfigServiceImpl{DB: db, Validator: validator, ConfigRepo: configRepo}
}

func (serv *ConfigServiceImpl) AddConfig(ctx context.Context, config model.Config) {
	rand.NewSource(time.Now().UnixNano())
	config.ID = helper.GenerateRandomString(15)

	txErr := serv.DB.Transaction(func(tx *gorm.DB) error {
		serv.ConfigRepo.AddConfig(ctx, tx, config)
		return nil
	})
	helper.Err(txErr)
}

func (serv *ConfigServiceImpl) GetConfigByBinId(ctx context.Context, binId string) web.ConfigGetResponse {
	var config model.Config
	txErr := serv.DB.Transaction(func(tx *gorm.DB) error {
		config = serv.ConfigRepo.GetCongifByBinId(ctx, tx, binId)
		return nil
	})
	helper.Err(txErr)
	return web.ConfigGetResponse{
		ConfigId:  config.ID,
		MaxHeight: config.MaxHeight,
		MaxWeight: config.MaxWeight,
	}
}

func (serv *ConfigServiceImpl) UpdateConfigById(ctx context.Context, request web.ConfigUpdateRequest, binId string) web.ConfigUpdateRespponse {
	errVal := serv.Validator.Struct(&request)
	helper.ValError(errVal)

	txErr := serv.DB.Transaction(func(tx *gorm.DB) error {
		serv.ConfigRepo.UpdateConfigById(ctx, tx, model.Config{
			BinID:     binId,
			MaxHeight: request.MaxHeight,
			MaxWeight: request.MaxWeight,
		})
		return nil
	})
	helper.Err(txErr)

	return web.ConfigUpdateRespponse{
		BinId:     binId,
		UpdatedAt: time.Now(),
	}
}
