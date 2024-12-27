package repository

import (
	"context"
	"smart-trash-bin/domain/model"

	"gorm.io/gorm"
)

type NotificationRepository interface {
	CreateNotification(ctx context.Context, tx *gorm.DB, notification model.Notification)
	UpdateNotificationById(ctx context.Context, tx *gorm.DB, notifId string, status bool)
	DeleteNotificaionById(ctx context.Context, tx *gorm.DB, notifId string)
}
