package service

import (
	"context"
	"smart-trash-bin/domain/web"
)

type UserService interface {
	Create(ctx context.Context, request web.UserCreateRequest) web.UserCreateResponse
	LoginUser(ctx context.Context, request web.UserLoginRequest) web.UserLoginResponse
	GetUserById(ctx context.Context, userId string) web.UserGetResponse
	UpdateUserById(ctx context.Context, request web.UserUpdateRequest, userId string) web.UserUpdateResponse
	UpdatePasswordById(ctx context.Context, request web.UserUpdatePasswordRequest, userId string) web.UserUpdateResponse
	LogoutUser(ctx context.Context, token string, userId string) web.UserLogoutResponse
}
