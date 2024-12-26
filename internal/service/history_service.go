package service

import (
	"context"
	"smart-trash-bin/domain/web"
)

type HistoryService interface {
	CreateHistory(ctx context.Context, request web.HistoryCreateRequest)
	GetHistoriesByBinId(ctx context.Context, binId string, userId string) web.HistoryGetResponse
}
