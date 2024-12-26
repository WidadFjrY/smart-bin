package repository

import (
	"context"
	"errors"
	"fmt"
	"smart-trash-bin/domain/model"
	"smart-trash-bin/pkg/exception"
	"smart-trash-bin/pkg/helper"
	"time"

	"gorm.io/gorm"
)

type GroupRepostoryImpl struct{}

func NewGroupRepository() GroupRepostory {
	return &GroupRepostoryImpl{}
}

func (repo *GroupRepostoryImpl) Create(ctx context.Context, tx *gorm.DB, group model.Group) model.Group {
	err := tx.WithContext(ctx).Create(&group).Error
	helper.Err(err)
	return group
}

func (repo *GroupRepostoryImpl) GetGroupById(ctx context.Context, tx *gorm.DB, groupId string) model.Group {
	var group model.Group
	err := tx.WithContext(ctx).Preload("SmartBin.Config").Where("id = ?", groupId).First(&group).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		panic(exception.NewNotFoundError(fmt.Sprintf("group with id %s not found", groupId)))
	}
	helper.Err(err)

	return group
}

func (repo *GroupRepostoryImpl) GetGroups(ctx context.Context, tx *gorm.DB, userId string, offset int, limit int) ([]model.Group, int64) {
	var group []model.Group
	var totalGroup int64

	err := tx.WithContext(ctx).Preload("SmartBin").Where("user_id = ?", userId).Find(&group).Offset(offset).Limit(limit).Error
	helper.Err(err)

	err = tx.WithContext(ctx).Table("groups").Where("user_id = ?", userId).Count(&totalGroup).Error
	helper.Err(err)

	return group, totalGroup
}

func (repo *GroupRepostoryImpl) UpdateGroupById(ctx context.Context, tx *gorm.DB, group model.Group) time.Time {
	err := tx.WithContext(ctx).Table("groups").Where("id = ?", group.ID).Updates(
		map[string]interface{}{
			"name":       group.Name,
			"location":   group.Location,
			"updated_at": time.Now(),
		},
	).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		panic(exception.NewNotFoundError(fmt.Sprintf("group with id %s not found", group.ID)))
	}
	helper.Err(err)

	return time.Now()
}
