package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"punishments-service/internal/http/service"
)

type PunishmentsHandler struct {
	services *service.Service
}

func NewPunishmentsHandler(services *service.Service) *PunishmentsHandler {
	return &PunishmentsHandler{
		services: services,
	}
}

func (h *PunishmentsHandler) GetPunishments(c *gin.Context) {
	limit, err := strconv.Atoi(c.Query("limit"))
	pType := c.Query("type")

	bans, err := h.services.GetPunishments(limit, pType)
	if err != nil {
		NewErrorResponse(c, http.StatusInternalServerError, fmt.Sprintf("Error receiving punishments: %s", err.Error()))

		return
	}

	if len(bans) == 0 {
		NewErrorResponse(c, http.StatusNotFound, "No punishments in the database yet")

		return
	}

	NewSuccessResponse(c, http.StatusOK, "Success receiving punishments", bans)
}
