package repository

import (
	"context"
	"smart-trash-bin/domain/model"
	"time"

	"gorm.io/gorm"
)

type ConfigRepository interface {
	AddConfig(ctx context.Context, tx *gorm.DB, config model.Config)
	GetCongifByBinId(ctx context.Context, tx *gorm.DB, binId string) model.Config
	UpdateConfigById(ctx context.Context, tx *gorm.DB, cofig model.Config) time.Time
}
