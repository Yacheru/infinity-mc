package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"news-service/init/logger"
	"news-service/internal/server/http/handlers"
	"news-service/pkg/constants"
	"strconv"
)

func ParseParam() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if _, err := strconv.Atoi(ctx.Param("id")); err != nil {
			logger.Error(err.Error(), constants.LoggerMiddlewares)
			handlers.NewErrorResponse(ctx, http.StatusBadRequest, constants.IdParamInvalid)
			return
		}

		ctx.Next()
	}
}
