package service

import (
	"context"
	"smart-trash-bin/domain/model"
	"smart-trash-bin/domain/web"
)

type ConfigService interface {
	AddConfig(ctx context.Context, config model.Config)
	GetConfigByBinId(ctx context.Context, binId string) web.ConfigGetResponse
	UpdateConfigById(ctx context.Context, request web.ConfigUpdateRequest, binId string) web.ConfigUpdateRespponse
}
