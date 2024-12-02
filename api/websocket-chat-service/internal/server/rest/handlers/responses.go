package handlers

import "github.com/gin-gonic/gin"

type Response struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"message"`
	Data       interface{} `json:"data,omitempty"`
}

func SuccessResponse(ctx *gin.Context, status int, message string, data interface{}) {
	ctx.AbortWithStatusJSON(status, Response{
		StatusCode: status,
		Message:    message,
		Data:       data,
	})
	return
}

func ErrorResponse(ctx *gin.Context, status int, message string) {
	ctx.AbortWithStatusJSON(status, Response{
		StatusCode: status,
		Message:    message,
	})
	return
}
