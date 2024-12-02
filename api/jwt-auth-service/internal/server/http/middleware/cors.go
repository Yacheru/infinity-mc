package middleware

import (
	"github.com/gin-gonic/gin"
	"jwt-auth-service/init/logger"
	"jwt-auth-service/internal/utils"
	"jwt-auth-service/pkg/constants"
	"net/http"
	"strings"
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
			logger.InfoF("method %s is not allowed\n", constants.MiddlewareCategory, ctx.Request.Method)
			ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden with CORS policy"})
			return
		}

		//for key, value := range ctx.Request.Header {
		//	if !utils.IsArrayContains(strings.Split(constants.AllowHeader, ", "), key) {
		//		logger.InfoF("init header %s: %s\n", constants.MiddlewareCategory, key, value)
		//		ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden with CORS policy"})
		//		return
		//	}
		//}

		if constants.AllowOrigin != "*" {
			if !utils.IsArrayContains(strings.Split(constants.AllowOrigin, ", "), "http://localhost:5173") {
				logger.InfoF("host '%s' is not part of '%v'\n", constants.MiddlewareCategory, ctx.Request.Host, constants.AllowOrigin)
				ctx.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden with CORS policy"})
				return
			}
		}
		
		ctx.Next()
	}
}
