package middleware

import (
	"smart-trash-bin/pkg/exception"

	"github.com/gin-gonic/gin"
)

func ErrorHandling() gin.HandlerFunc {
	return func(gin *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				exception.ErrorHandle(gin, err)
			}
		}()
		gin.Next()
	}
}
