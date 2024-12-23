package controller

import (
	"encoding/json"
	"net/http"
	"smart-trash-bin/domain/web"
	"smart-trash-bin/internal/service"
	"smart-trash-bin/pkg/exception"
	"smart-trash-bin/pkg/helper"

	"github.com/gin-gonic/gin"
)

type ConfigControllerImpl struct {
	ConfigServ service.ConfigService
}

func NewConfigController(configServ service.ConfigService) ConfigController {
	return &ConfigControllerImpl{ConfigServ: configServ}
}

func (cntrl *ConfigControllerImpl) UpdateConfigById(ctx *gin.Context) {
	var request web.ConfigUpdateRequest

	if ctx.Request.ContentLength == 0 {
		panic(exception.NewBadRequestError("request body required"))
	}

	binId := ctx.Params.ByName("bin_id")

	decoder := json.NewDecoder(ctx.Request.Body)
	err := decoder.Decode(&request)
	helper.Err(err)

	response := cntrl.ConfigServ.UpdateConfigById(ctx, request, binId)
	helper.Response(ctx, http.StatusOK, "Ok", response)
}
