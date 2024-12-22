package exception

import (
	"net/http"
	"smart-trash-bin/domain/web"

	"github.com/gin-gonic/gin"
)

func ErrorHandle(gin *gin.Context, err interface{}) {
	if badRequestError(gin, err) {
		return
	} else if notFoundError(gin, err) {
		return
	} else if unauthorized(gin, err) {
		return
	}

	internalServerError(gin, err)
}

func unauthorized(gin *gin.Context, err interface{}) bool {
	exception, ok := err.(Unauthorized)
	if ok {
		gin.JSON(http.StatusUnauthorized, web.ErrorResponse{
			Code:    http.StatusUnauthorized,
			Status:  "Unauthorized",
			Message: exception.error,
		})
		return true
	}
	return false
}

func notFoundError(gin *gin.Context, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		gin.JSON(http.StatusNotFound, web.ErrorResponse{
			Code:    http.StatusNotFound,
			Status:  "Not Found",
			Message: exception.error,
		})
		return true
	}
	return false
}

func badRequestError(gin *gin.Context, err interface{}) bool {
	exception, ok := err.(BadRequestError)
	if ok {
		gin.JSON(http.StatusBadRequest, web.ErrorResponse{
			Code:    http.StatusBadRequest,
			Status:  "Bad Request",
			Message: exception.error,
		})
		return true
	}
	return false
}

func internalServerError(gin *gin.Context, err interface{}) {
	gin.JSON(http.StatusInternalServerError, web.ErrorResponse{
		Code:    http.StatusInternalServerError,
		Status:  "Internal Server Error",
		Message: err.(error).Error(),
	})
}
