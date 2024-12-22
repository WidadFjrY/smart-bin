package repository

import (
	"context"
	"smart-trash-bin/domain/model"

	"gorm.io/gorm"
)

type ConfigRepository interface {
	AddConfig(ctx context.Context, tx *gorm.DB, config model.Config)
	GetCongifByBinId(ctx context.Context, tx *gorm.DB, binId string) model.Config
}
