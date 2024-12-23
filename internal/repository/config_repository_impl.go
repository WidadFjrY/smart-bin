package repository

import (
	"context"
	"errors"
	"smart-trash-bin/domain/model"
	"smart-trash-bin/pkg/exception"
	"smart-trash-bin/pkg/helper"
	"time"

	"gorm.io/gorm"
)

type ConfigRepositoryImpl struct{}

func NewConfigRepository() ConfigRepository {
	return &ConfigRepositoryImpl{}
}

func (repo *ConfigRepositoryImpl) AddConfig(ctx context.Context, tx *gorm.DB, config model.Config) {
	err := tx.WithContext(ctx).Create(&config).Error
	helper.Err(err)
}

func (repo *ConfigRepositoryImpl) GetCongifByBinId(ctx context.Context, tx *gorm.DB, binId string) model.Config {
	var config model.Config
	err := tx.WithContext(ctx).Where("bin_id = ?", binId).First(&config).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		panic(exception.NewNotFoundError("config not found"))
	}
	helper.Err(err)
	return config
}

func (repo *ConfigRepositoryImpl) UpdateConfigById(ctx context.Context, tx *gorm.DB, config model.Config) time.Time {
	err := tx.WithContext(ctx).Table("configs").Where("bin_id = ?", config.BinID).Updates(map[string]interface{}{
		"max_height": config.MaxHeight,
		"max_weight": config.MaxWeight,
		"updated_at": time.Now(),
	}).Error
	helper.Err(err)
	return time.Now()
}
