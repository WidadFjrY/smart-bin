package controller

import (
	"encoding/json"
	"net/http"
	"smart-trash-bin/domain/web"
	"smart-trash-bin/internal/service"
	"smart-trash-bin/pkg/exception"
	"smart-trash-bin/pkg/helper"
	"strings"

	"github.com/gin-gonic/gin"
)

type UserControllerImpl struct {
	UserServ service.UserService
}

func NewController(userServ service.UserService) UserController {
	return &UserControllerImpl{UserServ: userServ}
}

func (cntrl UserControllerImpl) Register(ctx *gin.Context) {
	var request web.UserCreateRequest

	if ctx.Request.ContentLength == 0 {
		panic(exception.NewBadRequestError("request body is empty"))
	}

	decoder := json.NewDecoder(ctx.Request.Body)
	err := decoder.Decode(&request)
	helper.Err(err)

	response := cntrl.UserServ.Create(ctx.Request.Context(), request)
	helper.Response(ctx, http.StatusCreated, "Created", response)
}

func (cntrl *UserControllerImpl) Login(ctx *gin.Context) {
	var request web.UserLoginRequest

	if ctx.Request.ContentLength == 0 {
		panic(exception.NewBadRequestError("request body is empty"))
	}

	decoder := json.NewDecoder(ctx.Request.Body)
	err := decoder.Decode(&request)
	helper.Err(err)

	response := cntrl.UserServ.LoginUser(ctx.Request.Context(), request)
	helper.Response(ctx, http.StatusOK, "OK", response)
}

func (cntrl *UserControllerImpl) GetUserById(ctx *gin.Context) {
	userId, _ := ctx.Get("user_id")

	response := cntrl.UserServ.GetUserById(ctx.Request.Context(), userId.(string))
	helper.Response(ctx, http.StatusOK, "Ok", response)
}

func (cntrl *UserControllerImpl) UpdateUserById(ctx *gin.Context) {
	var request web.UserUpdateRequest
	userId, _ := ctx.Get("user_id")

	if ctx.Request.ContentLength == 0 {
		panic(exception.NewBadRequestError("request body is empty"))
	}

	decoder := json.NewDecoder(ctx.Request.Body)
	err := decoder.Decode(&request)
	helper.Err(err)

	response := cntrl.UserServ.UpdateUserById(ctx.Request.Context(), request, userId.(string))
	helper.Response(ctx, http.StatusOK, "Ok", response)
}

func (cntrl *UserControllerImpl) UpdatePasswordById(ctx *gin.Context) {
	var request web.UserUpdatePasswordRequest
	userId, _ := ctx.Get("user_id")

	if ctx.Request.ContentLength == 0 {
		panic(exception.NewBadRequestError("request body is empty"))
	}

	decoder := json.NewDecoder(ctx.Request.Body)
	err := decoder.Decode(&request)
	helper.Err(err)

	response := cntrl.UserServ.UpdatePasswordById(ctx.Request.Context(), request, userId.(string))
	helper.Response(ctx, http.StatusOK, "Ok", response)
}

func (cntrl *UserControllerImpl) LogoutUser(ctx *gin.Context) {
	token := strings.Split(ctx.GetHeader("Authorization"), " ")[1]
	userId, _ := ctx.Get("user_id")

	response := cntrl.UserServ.LogoutUser(ctx.Request.Context(), token, userId.(string))
	helper.Response(ctx, http.StatusOK, "Ok", response)
}
