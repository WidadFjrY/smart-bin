package helper

import (
	"smart-trash-bin/domain/web"

	"github.com/gin-gonic/gin"
)

func Response(ctx *gin.Context, code int, status string, response interface{}) {
	ctx.JSON(code, web.SuccessResponse{
		Code:   code,
		Status: status,
		Data:   response,
	})
}

func ResponseWithPage(ctx *gin.Context, code int, status string, response interface{}, page int, totalPage int, totalItem int64) {
	ctx.JSON(code, web.SuccessResponseWithPage{
		Code:       code,
		Status:     status,
		Data:       response,
		Page:       page,
		TotalPages: totalPage,
		TotalItems: totalItem,
	})
}
