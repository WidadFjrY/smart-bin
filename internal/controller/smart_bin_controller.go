package controller

import "github.com/gin-gonic/gin"

type SmartBinController interface {
	AddSmartBin(ctx *gin.Context)
	UpdateSmartBinById(ctx *gin.Context)
	GetSmartBinById(ctx *gin.Context)
	DeleteSmartBinById(ctx *gin.Context)
	GetSmartBins(ctx *gin.Context)
	LockSmartBin(ctx *gin.Context)
	UnlockSmartBin(ctx *gin.Context)
	ClassifyImage(ctx *gin.Context)
	UpdateSmartBinValue(ctx *gin.Context)
}
