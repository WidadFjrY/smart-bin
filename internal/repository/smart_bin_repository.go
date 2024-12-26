package repository

import (
	"context"
	"smart-trash-bin/domain/model"
	"time"

	"gorm.io/gorm"
)

type SmartBinRepository interface {
	CreateSmartBin(ctx context.Context, tx *gorm.DB, smartBin model.SmartBin) model.SmartBin
	GetSmartBinById(ctx context.Context, tx *gorm.DB, binId string) (model.SmartBin, bool)
	UpdateSmartBinById(ctx context.Context, tx *gorm.DB, smartBin model.SmartBin) model.SmartBin
	DeleteSmartBinById(ctx context.Context, tx *gorm.DB, binId string) model.SmartBin
	GetSmartBins(ctx context.Context, tx *gorm.DB, limit int, offset int, userId string) []model.SmartBin
	TotalSmartBin(ctx context.Context, tx *gorm.DB, userId string) int64
	LockAndUnlockSmartBin(ctx context.Context, tx *gorm.DB, binId string, status bool) time.Time
	UpdateSensorValue(ctx context.Context, tx *gorm.DB, sensorValues map[string]float64, binId string) time.Time
	UpdateClassifyValue(ctx context.Context, tx *gorm.DB, waste map[string]float64, binId string) time.Time
	AddSmartBinToGroup(ctx context.Context, tx *gorm.DB, groupId string, binId string) time.Time
	RemoveSmartBinFromGroup(ctx context.Context, tx *gorm.DB, binId string) time.Time
	LockAndUnlockByGroup(ctx context.Context, tx *gorm.DB, groupId string, status bool) time.Time
}
