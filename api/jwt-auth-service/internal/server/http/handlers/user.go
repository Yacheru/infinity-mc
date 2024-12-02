package handlers

import (
	"errors"
	"jwt-auth-service/init/logger"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"jwt-auth-service/internal/entities"
	"jwt-auth-service/pkg/constants"
)

func (h *Handlers) SendCode(ctx *gin.Context) {
	var user = new(entities.SendCode)
	if err := ctx.ShouldBindJSON(user); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, "body is invalid")
		return
	}

	if err := h.s.UserService.SendCode(ctx.Request.Context(), user.Email); err != nil {
		if errors.Is(err, constants.UserAlreadyExistsError) {
			NewErrorResponse(ctx, http.StatusConflict, "User already exists")
			return
		}

		NewErrorResponse(ctx, http.StatusInternalServerError, "internal server error")
		return
	}

	NewSuccessResponse(ctx, http.StatusOK, "code sent", nil)
	return
}

// Register
// @Summary User SignUp
// @Tags user-auth
// @Description register account
// @Accept  json
// @Produce  json
// @Param input body entities.User true "sign up info"
// @Success 201 {object} string "ok"
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /auth/register [post]
func (h *Handlers) Register(ctx *gin.Context) {
	var user = new(entities.User)
	if err := ctx.ShouldBindJSON(user); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, "body is invalid")
		return
	}

	user.IpAddr = ctx.ClientIP()
	if strings.Contains(ctx.ClientIP(), "::1") {
		user.IpAddr = "127.0.0.1"
	}

	user.UserID = uuid.NewString()
	code, _ := strconv.Atoi(ctx.Query("code"))

	userResponse, err := h.s.UserService.Register(ctx.Request.Context(), user, code)
	if err != nil {
		if errors.Is(err, constants.CodeNotFoundError) {
			NewErrorResponse(ctx, http.StatusNotFound, "code not found")
			return
		}
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.SetCookie("refreshToken", userResponse.Tokens.RefreshToken, 30*24*60*60, "/", "", false, true)

	NewSuccessResponse(ctx, http.StatusCreated, "register success", userResponse)
	return
}

// Login
// @Summary User SignIn
// @Tags user-auth
// @Description login user
// @Accept  json
// @Produce  json
// @Param input body entities.UserLogin true "sign in info"
// @Success 201 {object} response
// @Failure 400,404 {object} response
// @Failure 500 {object} response
// @Failure default {object} response
// @Router /auth/login [post]
func (h *Handlers) Login(ctx *gin.Context) {
	var userLogin = new(entities.UserLogin)
	if err := ctx.ShouldBindJSON(userLogin); err != nil {
		NewErrorResponse(ctx, http.StatusBadRequest, "invalid body")
		return
	}

	userLogin.IpAddress = ctx.ClientIP()

	userResponse, err := h.s.UserService.Login(ctx.Request.Context(), userLogin)
	if err != nil {
		if errors.Is(err, constants.UserNotFoundError) {
			NewErrorResponse(ctx, http.StatusNotFound, "user not found")
			return
		}
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	ctx.SetCookie("refreshToken", userResponse.Tokens.RefreshToken, 30*24*60*60, "/", "", false, true)

	NewSuccessResponse(ctx, http.StatusOK, "client tokens", userResponse)
	return
}

func (h *Handlers) Logout(ctx *gin.Context) {
	logger.Debug("received logout request", ctx.ClientIP())

	refreshToken, err := ctx.Cookie("refreshToken")
	if err != nil {
		if errors.Is(err, http.ErrNoCookie) {
			NewErrorResponse(ctx, http.StatusUnauthorized, "no refresh token cookie")
			return
		}
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	logger.DebugF("cookie token received (%s)", ctx.ClientIP(), refreshToken)

	if err := h.s.UserService.Logout(ctx.Request.Context(), refreshToken); err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	logger.Debug("success response from service", ctx.ClientIP())

	http.SetCookie(ctx.Writer, &http.Cookie{
		Name:     "refreshToken",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		Expires:  time.Unix(0, 0),
		Secure:   false,
		HttpOnly: true,
	})

	logger.Debug("refreshToken deleted", ctx.ClientIP())

	NewSuccessResponse(ctx, http.StatusOK, "logout success", nil)
	return
}

func (h *Handlers) GetAll(ctx *gin.Context) {
	payload, exists := ctx.Get("payload")
	if !exists {
		NewErrorResponse(ctx, http.StatusBadRequest, "payload is missing")
		return
	}

	claims := payload.(*entities.Claims)
	if claims.Role != constants.Admin {
		NewErrorResponse(ctx, http.StatusForbidden, "forbidden")
		return
	}

	users, err := h.s.UserService.GetAll(ctx.Request.Context())
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	NewSuccessResponse(ctx, http.StatusOK, "users", users)
	return
}

func (h *Handlers) DeleteUser(ctx *gin.Context) {
	if err := h.s.UserService.DeleteUser(ctx.Request.Context(), ctx.Param("uuid")); err != nil {
		if errors.Is(err, constants.UserNotFoundError) {
			NewErrorResponse(ctx, http.StatusNotFound, "user not found")
			return
		}
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	NewSuccessResponse(ctx, http.StatusOK, "user deleted", nil)
	return
}

func (h *Handlers) UpdateRole(ctx *gin.Context) {
	if err := h.s.UserService.UpdateRole(ctx.Request.Context(), ctx.Param("uuid"), ctx.Query("role")); err != nil {
		if errors.Is(err, constants.UserNotFoundError) {
			NewErrorResponse(ctx, http.StatusNotFound, "user not found")
			return
		}
		NewErrorResponse(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	NewSuccessResponse(ctx, http.StatusOK, "role updated", nil)
	return
}
