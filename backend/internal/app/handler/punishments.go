package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/yacheru/infinity-mc.ru/backend/internal/app/repository"
	"github.com/yacheru/infinity-mc.ru/backend/internal/lib/api/response"
	"net/http"
	"strconv"
)

func (h *Handler) GetPunishments(c *gin.Context) {
	limit, err := strconv.Atoi(c.Query("limit"))
	pType := c.Query("type")

	pTypes := map[string]string{
		"bans":  repository.LbBans,
		"warns": repository.LbWarns,
		"mutes": repository.LbMutes,
	}
	pType, _ = pTypes[pType]

	bans, err := h.services.McBans.GetPunishments(limit, pType)
	if err != nil {
		logrus.Errorf("can't get punishments, %s", err.Error())

		response.NewErrorResponse(c, http.StatusInternalServerError, "can't get punishments", err.Error())

		return
	}
	if len(bans) == 0 {
		response.NewErrorResponse(c, http.StatusNoContent, "No punishments in the database yet", "")

		return
	}

	c.JSON(http.StatusOK, bans)
	return
}
