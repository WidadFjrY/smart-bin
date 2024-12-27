package service

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"smart-trash-bin/domain/model"
	"smart-trash-bin/domain/web"
	"smart-trash-bin/internal/repository"
	"smart-trash-bin/pkg/exception"
	"smart-trash-bin/pkg/helper"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserServiceImpl struct {
	DB        *gorm.DB
	Validator *validator.Validate
	UserRepo  repository.UserRepostiory
}

func NewUserService(db *gorm.DB, validator *validator.Validate, userRepo repository.UserRepostiory) UserService {
	return &UserServiceImpl{DB: db, Validator: validator, UserRepo: userRepo}
}

func (serv *UserServiceImpl) Create(ctx context.Context, request web.UserCreateRequest) web.UserCreateResponse {
	valErr := serv.Validator.Struct(&request)
	helper.ValError(valErr)

	if request.Password != request.VerifyPassword {
		panic(exception.NewBadRequestError("password doesn't match"))
	}

	hasedPass, err := bcrypt.GenerateFromPassword([]byte(request.Password), 12)
	helper.Err(err)

	rand.NewSource(time.Now().UnixNano())

	var userId string
	isExsit := true

	for isExsit {
		userId = helper.GenerateRandomString(15)
		txErr := serv.DB.Transaction(func(tx *gorm.DB) error {
			isExsit = !serv.UserRepo.IsUserExsit(ctx, tx, userId)
			return nil
		})
		helper.Err(txErr)
	}

	userName := fmt.Sprintf("User_%s", helper.GenerateRandomString(9))

	user := model.User{
		ID:       userId,
		Email:    request.Email,
		Password: string(hasedPass),
		Name:     userName,
	}

	var userResult model.User
	txErr := serv.DB.Transaction(func(tx *gorm.DB) error {
		userResult = serv.UserRepo.CreateUser(ctx, tx, user)
		return nil
	})
	helper.Err(txErr)

	return web.UserCreateResponse{
		Email:     userResult.Email,
		Name:      userResult.Name,
		CreatedAt: userResult.CreatedAt,
	}
}

func (serv *UserServiceImpl) LoginUser(ctx context.Context, request web.UserLoginRequest) web.UserLoginResponse {
	godotenv.Load()

	tokenRepo := repository.NewTokenRepository()
	tokenServ := NewTokenService(serv.DB, tokenRepo)

	valErr := serv.Validator.Struct(&request)
	helper.ValError(valErr)

	var user model.User
	txErr := serv.DB.Transaction(func(tx *gorm.DB) error {
		user = serv.UserRepo.GetUserByEmail(ctx, tx, request.Email)
		return nil
	})
	helper.Err(txErr)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if err != nil {
		panic(exception.NewUnauthorized("email or password wrong!"))
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, web.JwtClaims{
		Email:  user.Email,
		UserId: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})

	tokenStr, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		panic(err.Error())
	}

	tokenServ.Create(ctx, tokenStr)

	return web.UserLoginResponse{
		Token: tokenStr,
	}
}

func (serv *UserServiceImpl) GetUserById(ctx context.Context, userId string) web.UserGetResponse {
	var user model.User
	txErr := serv.DB.Transaction(func(tx *gorm.DB) error {
		user = serv.UserRepo.GetUserById(ctx, tx, userId)
		return nil
	})
	helper.Err(txErr)

	var bins []web.UserWithSmartBin
	var groups []web.UserWithGroup
	var notifications []web.UserWithNotif

	for _, smartBin := range user.SmartBin {
		bins = append(bins, web.UserWithSmartBin{
			BinId: smartBin.ID,
			Name:  smartBin.Name,
		})
	}

	for _, group := range user.Group {
		groups = append(groups, web.UserWithGroup{
			GroupId:   group.ID,
			Name:      group.Name,
			Location:  group.Location,
			TotalBins: len(group.SmartBin),
		})
	}

	for _, notification := range user.Notification {
		notifications = append(notifications, web.UserWithNotif{
			NotifId:   notification.ID,
			Title:     notification.Title,
			Desc:      notification.Desc,
			IsRead:    notification.IsRead,
			CreatedAt: notification.CreatedAt,
		})
	}

	return web.UserGetResponse{
		ID:           user.ID,
		Name:         user.Name,
		Email:        user.Email,
		SmartBin:     bins,
		Group:        groups,
		Notification: notifications,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}
}

func (serv *UserServiceImpl) UpdateUserById(ctx context.Context, request web.UserUpdateRequest, userId string) web.UserUpdateResponse {
	valErr := serv.Validator.Struct(&request)
	helper.ValError(valErr)

	var isExsit bool
	txErr := serv.DB.Transaction(func(tx *gorm.DB) error {
		isExsit = serv.UserRepo.IsUserExsit(ctx, tx, userId)
		return nil
	})
	helper.Err(txErr)
	if isExsit {
		panic(exception.NewNotFoundError(fmt.Sprintf("user with id %s not found", userId)))
	}

	userModel := model.User{
		ID:   userId,
		Name: request.Name,
	}

	var user model.User

	txErr = serv.DB.Transaction(func(tx *gorm.DB) error {
		user = serv.UserRepo.UpdateUserById(ctx, tx, userModel)
		return nil
	})
	helper.Err(txErr)

	return web.UserUpdateResponse{
		Id:        userId,
		UpdatedAt: user.UpdatedAt,
	}
}

func (serv *UserServiceImpl) UpdatePasswordById(ctx context.Context, request web.UserUpdatePasswordRequest, userId string) web.UserUpdateResponse {
	valErr := serv.Validator.Struct(&request)
	helper.ValError(valErr)

	var modelUser model.User
	serv.DB.Transaction(func(tx *gorm.DB) error {
		modelUser = serv.UserRepo.GetUserById(ctx, tx, userId)
		return nil
	})

	err := bcrypt.CompareHashAndPassword([]byte(modelUser.Password), []byte(request.Password))
	if err != nil {
		panic(exception.NewBadRequestError("password wrong!"))
	}

	if request.NewPassword != request.VerifyNewPassword {
		panic(exception.NewBadRequestError("password doesn't match!"))
	}

	hasedPassword, err := bcrypt.GenerateFromPassword([]byte(request.NewPassword), 12)
	helper.Err(err)

	modelUser = model.User{
		ID:       userId,
		Password: string(hasedPassword),
	}

	var user model.User
	txErr := serv.DB.Transaction(func(tx *gorm.DB) error {
		user = serv.UserRepo.UpdatePasswordById(ctx, tx, modelUser)
		return nil
	})
	helper.Err(txErr)

	return web.UserUpdateResponse{
		Id:        user.ID,
		UpdatedAt: user.UpdatedAt,
	}
}

func (serv *UserServiceImpl) LogoutUser(ctx context.Context, token string, userId string) web.UserLogoutResponse {
	tokenRepo := repository.NewTokenRepository()
	tokenServ := NewTokenService(serv.DB, tokenRepo)

	tokenServ.Update(ctx, token)

	return web.UserLogoutResponse{
		Id:          userId,
		LoggedOutAt: time.Now(),
	}
}
