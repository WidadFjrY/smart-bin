package controller

import (
	"net/http"
	"smart-trash-bin/domain/web"
	"smart-trash-bin/internal/service"
	"smart-trash-bin/pkg/helper"
	"time"

	"github.com/gin-gonic/gin"
)

type NotificationControllerImpl struct {
	Serv service.NotificationService
}

func NewNotificaitonController(serv service.NotificationService) NotificationController {
	return &NotificationControllerImpl{Serv: serv}
}

func (cntrl *NotificationControllerImpl) UpdateNotificationById(ctx *gin.Context) {
	notifId := ctx.Params.ByName("notif_id")
	cntrl.Serv.UpdateNotificationById(ctx.Request.Context(), notifId)
	helper.Response(ctx, http.StatusOK, "Ok", web.SuccessResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data: gin.H{
			"updated_at": time.Now(),
		},
	})
}

func (cntrl *NotificationControllerImpl) DeleteNotificaionById(ctx *gin.Context) {
	notifId := ctx.Params.ByName("notif_id")
	cntrl.Serv.DeleteNotificaionById(ctx.Request.Context(), notifId)
	helper.Response(ctx, http.StatusOK, "Ok", web.SuccessResponse{
		Code:   http.StatusOK,
		Status: "Ok",
		Data: gin.H{
			"updated_at": time.Now(),
		},
	})
}
