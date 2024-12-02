package handlers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"jwt-auth-service/init/logger"
	"jwt-auth-service/pkg/constants"
	"net/http"
)

// RefreshTokens
// @Summary User RefreshTokens
// @Tags tokens
// @Description sign-in user
// @Accept  json
// @Produce  json
// @Param input body entities.RefreshToken true "refresh tokens"
// @Param        guid    query     string  true  "User ID"
// @Success 200 {object} response
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /auth/refresh [post]
func (h *Handlers) RefreshTokens(ctx *gin.Context) {
	logger.Debug("received refresh token request", ctx.ClientIP())

	refreshToken, err := ctx.Cookie("refreshToken")
	if err != nil || refreshToken == "" {
		if errors.Is(err, http.ErrNoCookie) {
			NewErrorResponse(ctx, http.StatusUnauthorized, "no refresh token cookie")
			return
		}
		NewErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	logger.DebugF("cookie token received (%s)", ctx.ClientIP(), refreshToken)

	userResponse, err := h.s.AuthService.RefreshTokens(ctx.Request.Context(), refreshToken, ctx.ClientIP())
	if err != nil {
		if errors.Is(err, constants.UnauthorizedError) {
			NewErrorResponse(ctx, http.StatusUnauthorized, "unauthorized")
			return
		}
		if errors.Is(err, constants.UserNotFoundError) {
			NewErrorResponse(ctx, http.StatusUnauthorized, err.Error())
			return
		}
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	logger.DebugF("response from service (%+v)", ctx.ClientIP(), userResponse)

	ctx.SetCookie("refreshToken", userResponse.Tokens.RefreshToken, 30*24*60*60, "/", "", false, true)

	logger.Debug("refreshToken set in cookie", ctx.ClientIP())

	NewSuccessResponse(ctx, http.StatusOK, "new tokens", userResponse)
	return
}
