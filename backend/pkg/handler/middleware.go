package handler

import (
	"github.com/gin-gonic/gin"
)

const (
	Password = "psw"
)

func (h *Handler) ProtectApi(c *gin.Context) {
	//header := c.GetHeader(Password)
	//if header == "" {
	//	newErrorResponse(c, http.StatusUnauthorized, "empty auth header")
	//	return
	//}
}
