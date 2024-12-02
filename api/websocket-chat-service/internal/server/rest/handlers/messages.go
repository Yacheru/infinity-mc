package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (h *Handler) GetPlayerMessages(ctx *gin.Context) {
	nickname := ctx.Param("nickname")

	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "50"))
	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))

	messages, err := h.scylla.GetPlayerMessages(ctx.Request.Context(), nickname, limit, offset)
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	if len(messages) == 0 {
		ErrorResponse(ctx, http.StatusNotFound, "Message not found")
		return
	}

	SuccessResponse(ctx, http.StatusOK, fmt.Sprintf("%s messages", nickname), messages)
	return
}

func (h *Handler) GetAllMessages(ctx *gin.Context) {
	messages, err := h.scylla.GetAllMessages(ctx.Request.Context())
	if err != nil {
		ErrorResponse(ctx, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	SuccessResponse(ctx, http.StatusOK, "OK", messages)
	return
}
