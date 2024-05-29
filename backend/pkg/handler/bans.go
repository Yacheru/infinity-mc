package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yacheru/infinity-mc.ru/backend/internal/lib/api/response"
	"net/http"
	"strconv"
)

func (h *Handler) GetAllBans(c *gin.Context) {
	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil {
		response.NewErrorResponse(c, http.StatusBadRequest, "invalid limit parameter", err.Error())
		return
	}

	bans, err := h.services.McBans.GetAllBans(limit)
	if err != nil {

		response.NewErrorResponse(c, http.StatusInternalServerError, "can't get all bans", err.Error())

		return
	}

	c.JSON(http.StatusOK, bans)
}
