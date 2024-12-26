package service

import (
	"context"
	"smart-trash-bin/domain/web"
)

type GroupService interface {
	Create(ctx context.Context, request web.GroupCreateRequest, userId string) web.GroupCreateResponse
	GetGroupById(ctx context.Context, groupId string, userId string) web.GroupGetResponse
	GetGroups(ctx context.Context, page int, userId string) ([]web.GroupGetResponses, int64, int)
}
