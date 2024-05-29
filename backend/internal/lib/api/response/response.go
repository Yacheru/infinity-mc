package response

import (
	"errors"
	"github.com/gin-gonic/gin"
)

type errorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Error   string `json:"description"`
}

func NewErrorResponse(c *gin.Context, statusCode int, message string, err string) {
	if err == "" {
		err = "not provided"
	}

	c.AbortWithStatusJSON(statusCode, errorResponse{statusCode, message, errors.New(err).Error()})
}
