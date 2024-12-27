package web

import (
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
	ID           string             `json:"id"`
	Name         string             `json:"name"`
	Email        string             `json:"email"`
	SmartBin     []UserWithSmartBin `json:"smart_bin"`
	Group        []UserWithGroup    `json:"group"`
	Notification []UserWithNotif
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type UserWithSmartBin struct {
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

type UserWithGroup struct {
	GroupId   string `json:"group_id"`
	Name      string `json:"name"`
	Location  string `json:"location"`
	TotalBins int    `json:"total_bins"`
}

type UserWithNotif struct {
	NotifId   string    `json:"notif_id"`
	Title     string    `json:"title"`
	Desc      string    `json:"desc"`
	IsRead    bool      `json:"is_read"`
	CreatedAt time.Time `json:"created_at"`
}
