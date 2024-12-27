package service

import (
	"context"
	"math/rand"
	"smart-trash-bin/domain/model"
	"smart-trash-bin/domain/web"
	"smart-trash-bin/internal/repository"
	"smart-trash-bin/pkg/helper"
	"time"

	"gorm.io/gorm"
)

type NotificationServiceImpl struct {
	DB   *gorm.DB
	Repo repository.NotificationRepository
}

func NewNotificationService(db *gorm.DB, repo repository.NotificationRepository) NotificationService {
	return &NotificationServiceImpl{DB: db, Repo: repo}
}

func (serv *NotificationServiceImpl) CreateNotification(ctx context.Context, request web.NotificationCreateRequest, userId string) {
	rand.NewSource(time.Now().Unix())

	txErr := serv.DB.Transaction(func(tx *gorm.DB) error {
		serv.Repo.CreateNotification(ctx, tx, model.Notification{
			ID:     helper.GenerateRandomString(15),
			UserID: userId,
			Title:  request.Title,
			Desc:   request.Desc,
		})
		return nil
	})
	helper.Err(txErr)

}

func (serv *NotificationServiceImpl) UpdateNotificationById(ctx context.Context, notifId string) {
	txErr := serv.DB.Transaction(func(tx *gorm.DB) error {
		serv.Repo.UpdateNotificationById(ctx, tx, notifId, true)
		return nil
	})
	helper.Err(txErr)
}

func (serv *NotificationServiceImpl) DeleteNotificaionById(ctx context.Context, notifId string) {
	txErr := serv.DB.Transaction(func(tx *gorm.DB) error {
		serv.Repo.DeleteNotificaionById(ctx, tx, notifId)
		return nil
	})
	helper.Err(txErr)
}
