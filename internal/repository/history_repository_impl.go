package repository

import (
	"context"
	"errors"
	"smart-trash-bin/domain/model"
	"smart-trash-bin/pkg/exception"
	"smart-trash-bin/pkg/helper"

	"gorm.io/gorm"
)

type HistoryRepositoryImpl struct{}

func NewHistoryRepository() HistoryRepository {
	return &HistoryRepositoryImpl{}
}

func (repo *HistoryRepositoryImpl) CreateHistory(ctx context.Context, tx *gorm.DB, history model.History) {
	err := tx.WithContext(ctx).Create(&history).Error
	helper.Err(err)
}

func (repo *HistoryRepositoryImpl) GetHistoriesByBinId(ctx context.Context, tx *gorm.DB, binId string) []model.History {
	var histories []model.History
	err := tx.WithContext(ctx).Where("bin_id = ?", binId).Find(&histories).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		panic(exception.NewBadRequestError("history empty"))
	}
	helper.Err(err)
	return histories
}
