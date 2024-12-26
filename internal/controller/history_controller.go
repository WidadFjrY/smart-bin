package controller

import "github.com/gin-gonic/gin"

type HistoryController interface {
	GetHistoriesByBinId(ctx *gin.Context)
}
