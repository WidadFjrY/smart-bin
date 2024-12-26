package repository

import (
	"context"
	"errors"
	"fmt"
	"smart-trash-bin/domain/model"
	"smart-trash-bin/pkg/exception"
	"smart-trash-bin/pkg/helper"
	"strings"
	"time"

	"gorm.io/gorm"
)

type SmartBinRepositoryImpl struct{}

func NewSmartBinRepository() SmartBinRepository {
	return &SmartBinRepositoryImpl{}
}

func (repo *SmartBinRepositoryImpl) CreateSmartBin(ctx context.Context, tx *gorm.DB, smartBin model.SmartBin) model.SmartBin {

	result := tx.WithContext(ctx).Create(&smartBin)
	if result.Error != nil {
		if strings.Contains(result.Error.Error(), "Error 1062") {
			smartBin.DeletedAt.Valid = false
			err := tx.WithContext(ctx).Save(&smartBin).Error
			helper.Err(err)
		} else {
			panic(result.Error.Error())
		}
	}

	smartBin.CreatedAt = time.Now()

	return smartBin
}

func (repo *SmartBinRepositoryImpl) GetSmartBinById(ctx context.Context, tx *gorm.DB, binId string) (model.SmartBin, bool) {
	var smartBin model.SmartBin

	err := tx.WithContext(ctx).Preload("Config").Where("id = ?", binId).First(&smartBin).Error
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return smartBin, true
	}

	return smartBin, false
}

func (repo *SmartBinRepositoryImpl) UpdateSmartBinById(ctx context.Context, tx *gorm.DB, smartBin model.SmartBin) model.SmartBin {
	err := tx.WithContext(ctx).Table("smart_bins").Where("id = ?", smartBin.ID).Updates(
		map[string]interface{}{
			"name":       smartBin.Name,
			"location":   smartBin.Location,
			"updated_at": time.Now(),
		}).Error
	helper.Err(err)

	smartBin.UpdatedAt = time.Now()
	return smartBin
}

func (repo *SmartBinRepositoryImpl) DeleteSmartBinById(ctx context.Context, tx *gorm.DB, binId string) model.SmartBin {
	err := tx.WithContext(ctx).Where("id = ?", binId).Delete(&model.SmartBin{}).Error
	helper.Err(err)

	return model.SmartBin{
		ID:        binId,
		DeletedAt: gorm.DeletedAt{Time: time.Now()},
	}
}

func (repo *SmartBinRepositoryImpl) GetSmartBins(ctx context.Context, tx *gorm.DB, limit int, offset int, userId string) []model.SmartBin {
	var smartBins []model.SmartBin

	err := tx.WithContext(ctx).Preload("Config").Where("user_id = ?", userId).Limit(limit).Offset(offset).Find(&smartBins).Error
	helper.Err(err)

	return smartBins
}

func (repo *SmartBinRepositoryImpl) TotalSmartBin(ctx context.Context, tx *gorm.DB, userId string) int64 {
	var totalItems int64
	err := tx.WithContext(ctx).Table("smart_bins").Where("user_id = ?", userId).Count(&totalItems).Error
	helper.Err(err)

	return totalItems
}

func (repo *SmartBinRepositoryImpl) LockAndUnlockSmartBin(ctx context.Context, tx *gorm.DB, binId string, status bool) time.Time {
	// lock is true, unlock is false
	err := tx.WithContext(ctx).Table("smart_bins").Where("id = ?", binId).Updates(
		map[string]interface{}{
			"is_locked":  status,
			"updated_at": time.Now(),
		},
	).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		panic(exception.NewNotFoundError(fmt.Sprintf("smart bin with id %s not found", binId)))
	}
	helper.Err(err)

	return time.Now()

}

func (repo *SmartBinRepositoryImpl) UpdateSensorValue(ctx context.Context, tx *gorm.DB, sensorValues map[string]float64, binId string) time.Time {
	err := tx.WithContext(ctx).Table("smart_bins").Where("id = ?", binId).Updates(
		map[string]interface{}{
			"load_cell_value": gorm.Expr(`
				JSON_SET(
					load_cell_value,
					'$.organic', ?,
					'$.non_organic', ?
				)`,
				sensorValues["load_cell_organic"], sensorValues["load_cell_non_organic"],
			),
			"ultra_sonic_value": gorm.Expr(`
				JSON_SET(
					ultra_sonic_value,
					'$.organic', ?,
					'$.non_organic', ?
				)`,
				sensorValues["ultra_sonic_organic"], sensorValues["ultra_sonic_non_organic"],
			),
			"updated_at": time.Now(),
		},
	).Error
	helper.Err(err)
	return time.Now()
}

func (repo *SmartBinRepositoryImpl) UpdateClassifyValue(ctx context.Context, tx *gorm.DB, waste map[string]float64, binId string) time.Time {
	err := tx.WithContext(ctx).Table("smart_bins").Where("id = ?", binId).Updates(
		map[string]interface{}{
			"total_organic_waste":     gorm.Expr("total_organic_waste + ?", waste["organic"]),
			"total_non_organic_waste": gorm.Expr("total_non_organic_waste + ?", waste["non_organic"]),
			"updated_at":              time.Now(),
		},
	).Error
	helper.Err(err)
	return time.Now()
}

func (repo *SmartBinRepositoryImpl) AddSmartBinToGroup(ctx context.Context, tx *gorm.DB, groupId string, binId string) time.Time {
	err := tx.WithContext(ctx).Table("smart_bins").Where("id = ?", binId).Updates(
		map[string]interface{}{
			"group_id":   groupId,
			"updated_at": time.Now(),
		},
	).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		panic(exception.NewNotFoundError(fmt.Sprintf("smart bin with id %s not found", binId)))
	}
	helper.Err(err)
	return time.Now()
}

func (repo *SmartBinRepositoryImpl) RemoveSmartBinFromGroup(ctx context.Context, tx *gorm.DB, binId string) time.Time {
	err := tx.WithContext(ctx).Table("smart_bins").Where("id = ?", binId).Updates(
		map[string]interface{}{
			"group_id":   nil,
			"updated_at": time.Now(),
		},
	).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		panic(exception.NewNotFoundError(fmt.Sprintf("smart bin with id %s not found", binId)))
	}
	helper.Err(err)
	return time.Now()
}
