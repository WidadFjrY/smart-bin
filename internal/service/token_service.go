package service

import "context"

type TokenService interface {
	Create(ctx context.Context, token string)
	IsValid(ctx context.Context, token string) bool
	Update(ctx context.Context, token string)
}
