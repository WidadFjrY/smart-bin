package service

import (
	"context"
	"smart-trash-bin/domain/web"
)

type SmartBinService interface {
	AddSmartBin(ctx context.Context, request web.SmartBinCreateRequest, userId string) web.SmartBinCreateResponse
	IsSmartBinExsist(ctx context.Context, binId string)
	UpdateSmartBinById(ctx context.Context, request web.SmartBinUpdateRequest, binId string, userId string) web.SmartBinUpdateResponse
	GetSmartBinById(ctx context.Context, binId string, userId string) web.SmartBinGetResponse
	DeleteSmartBinById(ctx context.Context, binId string, userId string) web.SmartBinDeleteResponse
	GetSmartBins(ctx context.Context, page int, userId string) ([]web.SmartBinGetResponse, int64, int)
	LockAndUnlockSmartBin(ctx context.Context, status bool, binId string) web.SmartBinUpdateResponse
	ClassifyImage(ctx context.Context, binId string, classify web.ClassifyResponse) web.SmartBinClassification
	UpdateDataSmartBin(ctx context.Context, binId string, request web.UpdateValueRequest) web.UpdateValueResponse
	IsSmartBinFull(ctx context.Context, status bool, binId string)
}
