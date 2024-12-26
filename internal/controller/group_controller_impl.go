package controller

import (
	"encoding/json"
	"net/http"
	"smart-trash-bin/domain/web"
	"smart-trash-bin/internal/service"
	"smart-trash-bin/pkg/exception"
	"smart-trash-bin/pkg/helper"
	"strconv"

	"github.com/gin-gonic/gin"
)

type GroupControllerImpl struct {
	Serv service.GroupService
}

func NewGroupController(serv service.GroupService) GroupController {
	return &GroupControllerImpl{Serv: serv}
}

func (cntrl *GroupControllerImpl) Create(ctx *gin.Context) {
	var request web.GroupCreateRequest
	userId, _ := ctx.Get("user_id")

	if ctx.Request.ContentLength == 0 {
		panic(exception.NewBadRequestError("request body required"))
	}

	decoder := json.NewDecoder(ctx.Request.Body)
	err := decoder.Decode(&request)
	helper.Err(err)

	response := cntrl.Serv.Create(ctx.Request.Context(), request, userId.(string))
	helper.Response(ctx, http.StatusCreated, "Created", response)
}

func (cntrl *GroupControllerImpl) GetGroupById(ctx *gin.Context) {
	userId, _ := ctx.Get("user_id")
	groupId := ctx.Params.ByName("group_id")

	response := cntrl.Serv.GetGroupById(ctx.Request.Context(), groupId, userId.(string))
	helper.Response(ctx, http.StatusOK, "Ok", response)
}

func (cntrl *GroupControllerImpl) GetGroups(ctx *gin.Context) {
	userId, _ := ctx.Get("user_id")
	page := ctx.Params.ByName("page")
	pageInt, err := strconv.Atoi(page)

	if err != nil {
		panic(exception.NewBadRequestError("page must be a number!"))
	}

	response, totalItems, totalPage := cntrl.Serv.GetGroups(ctx.Request.Context(), pageInt, userId.(string))
	helper.ResponseWithPage(ctx, http.StatusOK, "Ok", response, pageInt, totalPage, totalItems)
}

func (cntrl *GroupControllerImpl) UpdateGroupById(ctx *gin.Context) {
	var request web.GroupUpdateRequest
	userId, _ := ctx.Get("user_id")
	groupId := ctx.Params.ByName("group_id")

	decoder := json.NewDecoder(ctx.Request.Body)
	decoder.Decode(&request)

	response := cntrl.Serv.UpdateGroupById(ctx.Request.Context(), request, groupId, userId.(string))
	helper.Response(ctx, http.StatusOK, "Ok", response)
}

func (cntrl *GroupControllerImpl) DeleteGroupById(ctx *gin.Context) {
	userId, _ := ctx.Get("user_id")
	groupId := ctx.Params.ByName("group_id")

	response := cntrl.Serv.DeleteGroupById(ctx.Request.Context(), groupId, userId.(string))
	helper.Response(ctx, http.StatusOK, "Ok", response)
}
