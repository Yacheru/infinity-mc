package middlewares

import (
	"net/http"
	"punishments-service/internal/utils"
	"punishments-service/pkg/constants"
	"strings"

	"github.com/gin-gonic/gin"
	"punishments-service/init/logger"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", constants.AllowOrigin)
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", constants.AllowCredential)
		ctx.Writer.Header().Set("Access-Control-Allow-Headers", constants.AllowHeader)
		ctx.Writer.Header().Set("Access-Control-Allow-Methods", constants.AllowMethods)
		ctx.Writer.Header().Set("Access-Control-Max-Age", constants.MaxAge)

		if ctx.Request.Method == http.MethodOptions {
			ctx.AbortWithStatus(http.StatusOK)
			return
		}

		if !utils.IsArrayContains(strings.Split(constants.AllowMethods, ", "), ctx.Request.Method) {
			logger.InfoF("method %s is not allowed\n", constants.LoggerCategoryCORS, ctx.Request.Method)
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden with CORS policy"})
			return
		}

		//for key, value := range ctx.Request.Header {
		//	if !utils.IsArrayContains(strings.Split(constants.AllowHeader, ", "), key) {
		//		logger.InfoF("init header %s: %s\n", logrus.Fields{constants.LoggerCategory: constants.LoggerCategoryCORS}, key, value)
		//		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden with CORS policy"})
		//		return
		//	}
		//}

		if constants.AllowOrigin != "*" {
			if !utils.IsArrayContains(strings.Split(constants.AllowOrigin, ", "), "http://localhost:5173") {
				logger.InfoF("host '%s' is not part of '%v'\n", constants.LoggerCategoryCORS, ctx.Request.Host, constants.AllowOrigin)
				ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden with CORS policy"})
				return
			}
		}

		ctx.Next()
	}
}
