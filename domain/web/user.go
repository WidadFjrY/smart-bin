package web

import (
	"smart-trash-bin/domain/model"
	"time"
)

type UserCreateRequest struct {
	Email          string `json:"email" validate:"required,email,min=1"`
	Password       string `json:"password" validate:"required,min=8"`
	VerifyPassword string `json:"verify_password" validate:"required,min=8"`
}

type UserCreateResponse struct {
	Email     string    `json:"email"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

type UserLoginRequest struct {
	Email    string `json:"email" validate:"required,email,min=1"`
	Password string `json:"password" validate:"required,min=8"`
}

type UserLoginResponse struct {
	Token string `json:"token"`
}

type UserGetResponse struct {
	ID        string                     `json:"id"`
	Name      string                     `json:"name"`
	Email     string                     `json:"email"`
	SmartBin  []SmartBinWithUserResponse `json:"smart_bin"`
	Group     []model.Group              `json:"group"`
	CreatedAt time.Time                  `json:"created_at"`
	UpdatedAt time.Time                  `json:"updated_at"`
}

type SmartBinWithUserResponse struct {
	BinId string `json:"bin_id"`
	Name  string `json:"name"`
}

type UserUpdateRequest struct {
	Name string `json:"name" validate:"required,min=3"`
}

type UserUpdatePasswordRequest struct {
	Password          string `json:"password" validate:"required,min=8"`
	NewPassword       string `json:"new_password" validate:"required,min=8"`
	VerifyNewPassword string `json:"verify_new_password" validate:"required,min=8"`
}

type UserUpdateResponse struct {
	Id        string    `json:"id"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserLogoutResponse struct {
	Id          string    `json:"id"`
	LoggedOutAt time.Time `json:"logged_out_at"`
}
