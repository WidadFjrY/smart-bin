package controller

import (
	"net/http"
	"smart-trash-bin/internal/service"
	"smart-trash-bin/pkg/helper"

	"github.com/gin-gonic/gin"
)

type HistoryControllerImpl struct {
	Serv service.HistoryService
}

func NewHistoryController(serv service.HistoryService) HistoryController {
	return &HistoryControllerImpl{Serv: serv}
}

func (cntrl *HistoryControllerImpl) GetHistoriesByBinId(ctx *gin.Context) {
	binId := ctx.Params.ByName("bin_id")
	userId, _ := ctx.Get("user_id")

	response := cntrl.Serv.GetHistoriesByBinId(ctx.Request.Context(), binId, userId.(string))
	helper.Response(ctx, http.StatusOK, "Ok", response)
}
