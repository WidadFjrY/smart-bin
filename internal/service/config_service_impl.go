package service

import (
	"context"
	"math/rand"
	"smart-trash-bin/domain/model"
	"smart-trash-bin/domain/web"
	"smart-trash-bin/internal/repository"
	"smart-trash-bin/pkg/helper"
	"time"

	"gorm.io/gorm"
)

type ConfigServiceImpl struct {
	DB         *gorm.DB
	ConfigRepo repository.ConfigRepository
}

func NewConfigService(db *gorm.DB, configRepo repository.ConfigRepository) ConfigService {
	return &ConfigServiceImpl{DB: db, ConfigRepo: configRepo}
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
		Id:        config.ID,
		MaxHeight: config.MaxHeight,
		MaxWeight: config.MaxWeight,
	}
}
