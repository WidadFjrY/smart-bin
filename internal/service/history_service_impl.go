package service

import (
	"context"
	"fmt"
	"math/rand"
	"smart-trash-bin/domain/model"
	"smart-trash-bin/domain/web"
	"smart-trash-bin/internal/repository"
	"smart-trash-bin/pkg/helper"
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type HistoryServiceImpl struct {
	DB   *gorm.DB
	Repo repository.HistoryRepository
}

func NewHistoryService(db *gorm.DB, repo repository.HistoryRepository) HistoryService {
	return &HistoryServiceImpl{DB: db, Repo: repo}
}

func (serv *HistoryServiceImpl) CreateHistory(ctx context.Context, request web.HistoryCreateRequest) {
	rand.NewSource(time.Now().Unix())
	historyId := helper.GenerateRandomString(15)

	fmt.Println(historyId)

	txErr := serv.DB.Transaction(func(tx *gorm.DB) error {
		serv.Repo.CreateHistory(ctx, tx, model.History{
			ID:      historyId,
			BinID:   request.BinID,
			Status:  request.Status,
			Message: request.Message,
		})
		return nil
	})

	helper.Err(txErr)
}

func (serv *HistoryServiceImpl) GetHistoriesByBinId(ctx context.Context, binId string, userId string) web.HistoryGetResponse {

	binRepo := repository.NewSmartBinRepository()
	binServ := NewSmartBinService(serv.DB, validator.New(), binRepo)

	binServ.GetSmartBinById(ctx, binId, userId)

	var histories []model.History

	errTx := serv.DB.Transaction(func(tx *gorm.DB) error {
		histories = serv.Repo.GetHistoriesByBinId(ctx, tx, binId)
		return nil
	})
	helper.Err(errTx)

	var historyResponse web.HistoryGetResponse
	for _, history := range histories {
		desc := web.History{
			Status:    history.Status,
			Message:   history.Message,
			CreatedAt: history.CreatedAt,
		}

		historyResponse.BinId = history.BinID
		historyResponse.History = append(historyResponse.History, desc)
	}

	return historyResponse
}
