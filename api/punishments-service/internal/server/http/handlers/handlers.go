package handlers

import (
	"fmt"
	"net/http"
	"punishments-service/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type PunishmentsHandler struct {
	services *service.Service
}

func NewPunishmentsHandler(services *service.Service) *PunishmentsHandler {
	return &PunishmentsHandler{
		services: services,
	}
}

func (h *PunishmentsHandler) GetPunishments(ctx *gin.Context) {
	limit, err := strconv.Atoi(ctx.Query("limit"))

	bans, err := h.services.GetPunishments(ctx.Request.Context(), limit, ctx.Query("type"))
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, fmt.Sprintf("Error receiving punishments: %s", err.Error()))
		return
	}

	if len(*bans) == 0 {
		NewErrorResponse(ctx, http.StatusNoContent, "No punishments in the database yet")
		return
	}

	NewSuccessResponse(ctx, http.StatusOK, "Success receiving punishments", bans)
	return
}

func (h *PunishmentsHandler) Unban(ctx *gin.Context) {

}
