package handlers

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func NewSuccessResponse(c *gin.Context, statusCode int, message string, data interface{}) {
	c.AbortWithStatusJSON(statusCode, Response{
		Status:  statusCode,
		Message: message,
		Data:    data,
	})
}

func NewErrorResponse(c *gin.Context, statusCode int, message string) {
	c.AbortWithStatusJSON(statusCode, Response{
		Status:  statusCode,
		Message: message,
	})
}
