package repository

import (
	"context"
	"smart-trash-bin/domain/model"

	"gorm.io/gorm"
)

type HistoryRepository interface {
	CreateHistory(ctx context.Context, tx *gorm.DB, history model.History)
	GetHistoriesByBinId(ctx context.Context, tx *gorm.DB, binId string) []model.History
}
