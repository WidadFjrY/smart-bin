package service

import (
	"context"
	"smart-trash-bin/domain/web"
)

type NotificationService interface {
	CreateNotification(ctx context.Context, request web.NotificationCreateRequest, userId string)
	UpdateNotificationById(ctx context.Context, notifId string)
	DeleteNotificaionById(ctx context.Context, notifId string)
}
