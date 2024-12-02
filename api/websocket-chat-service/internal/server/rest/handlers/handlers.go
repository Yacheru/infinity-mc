package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"websocket-chat-service/internal/server/websocket"
	"websocket-chat-service/internal/service"
	"websocket-chat-service/pkg/constants"
)

type Handler struct {
	scylla    service.ScyllaService
	websocket websocket.Client
	wg        sync.WaitGroup
}

func NewHandler(scylla service.ScyllaService, websocket websocket.Client) *Handler {
	return &Handler{
		scylla:    scylla,
		websocket: websocket,
		wg:        sync.WaitGroup{},
	}
}

func (h *Handler) CloseWS(ctx *gin.Context) {
	errorsMap := h.websocket.CloseConnections()
	if len(errorsMap) > 0 {
		ctx.AbortWithStatusJSON(http.StatusServiceUnavailable, Response{
			StatusCode: http.StatusInternalServerError,
			Message:    "websocket connections error",
			Data:       errorsMap,
		})
		return
	}

	SuccessResponse(ctx, http.StatusOK, "All connections closed successfully", nil)
	return
}

func (h *Handler) RunWS(ctx *gin.Context) {
	h.wg.Add(1)
	if err := h.websocket.Dial(ctx); err != nil {
		h.wg.Done()
		if errors.Is(err, constants.MaxLimitConnError) {
			ErrorResponse(ctx, http.StatusConflict, err.Error())
			return
		}

		ErrorResponse(ctx, http.StatusInternalServerError, "Internal server error")
		return
	}
	h.wg.Done()

	SuccessResponse(ctx, http.StatusOK, "Success", nil)
	return
}
