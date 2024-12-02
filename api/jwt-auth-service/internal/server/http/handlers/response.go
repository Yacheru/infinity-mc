package handlers

import (
	"github.com/gin-gonic/gin"
)

type response struct {
	StatusCode int         `json:"status"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

func NewSuccessResponse(ctx *gin.Context, status int, message string, data interface{}) {
	ctx.AbortWithStatusJSON(status, response{
		StatusCode: status,
		Message:    message,
		Data:       data,
	})
}

func NewErrorResponse(ctx *gin.Context, status int, message string) {
	ctx.AbortWithStatusJSON(status, response{
		StatusCode: status,
		Message:    message,
	})
}
