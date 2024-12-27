package repository

import (
	"context"
	"smart-trash-bin/domain/model"
	"smart-trash-bin/pkg/helper"

	"gorm.io/gorm"
)

type NotificationRepositoryImpl struct{}

func NewNotificationRepository() NotificationRepository {
	return &NotificationRepositoryImpl{}
}

func (repo *NotificationRepositoryImpl) CreateNotification(ctx context.Context, tx *gorm.DB, notification model.Notification) {
	err := tx.WithContext(ctx).Create(&notification).Error
	helper.Err(err)
}

func (repo *NotificationRepositoryImpl) UpdateNotificationById(ctx context.Context, tx *gorm.DB, notifId string, status bool) {
	err := tx.WithContext(ctx).Table("notifications").Where("id = ?", notifId).Updates(
		map[string]interface{}{
			"is_read": status,
		},
	).Error
	helper.Err(err)
}

func (repo *NotificationRepositoryImpl) DeleteNotificaionById(ctx context.Context, tx *gorm.DB, notifId string) {
	err := tx.WithContext(ctx).Table("notifications").Where("id = ?", notifId).Delete(nil).Error
	helper.Err(err)
}
