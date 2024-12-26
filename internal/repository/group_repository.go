package repository

import (
	"context"
	"smart-trash-bin/domain/model"
	"time"

	"gorm.io/gorm"
)

type GroupRepostory interface {
	Create(ctx context.Context, tx *gorm.DB, group model.Group) model.Group
	GetGroupById(ctx context.Context, tx *gorm.DB, groupId string) model.Group
	GetGroups(ctx context.Context, tx *gorm.DB, userId string, offset int, limit int) ([]model.Group, int64)
	UpdateGroupById(ctx context.Context, tx *gorm.DB, group model.Group) time.Time
}
