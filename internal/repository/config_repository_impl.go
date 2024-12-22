package repository

import (
	"context"
	"errors"
	"smart-trash-bin/domain/model"
	"smart-trash-bin/pkg/exception"
	"smart-trash-bin/pkg/helper"

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
