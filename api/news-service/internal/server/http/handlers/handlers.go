package handlers

import (
	"errors"
	"net/http"
	"news-service/internal/service"
	"strconv"

	"github.com/gin-gonic/gin"

	"news-service/init/logger"
	"news-service/internal/entities"
	"news-service/pkg/constants"
)

type Handlers struct {
	service service.News
}

func NewHandlers(service *service.Service) *Handlers {
	return &Handlers{service: service}
}

func (h *Handlers) AddNew(ctx *gin.Context) {
	var embedEntity = new(entities.Embed)
	if err := ctx.ShouldBindJSON(embedEntity); err != nil {
		logger.Error(err.Error(), constants.LoggerHandlers)
		NewErrorResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.service.AddNew(ctx, embedEntity)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	NewSuccessResponse(ctx, http.StatusOK, "news added successfully", id)
	return
}

func (h *Handlers) GetAllNews(ctx *gin.Context) {
	// TODO:
	/*
		- Search by discord-id
		- Search by title
		- Search by description
		- Search by author
		- Search by footer
	*/
	//entityParams, exists := ctx.Get("params")
	//if !exists {
	//	logger.Error(constants.ParamsDoesntExistsError.Error(), constants.LoggerHandlers)
	//	NewErrorResponse(ctx, http.StatusInternalServerError, constants.ParamsDoesntExistsError.Error())
	//	return
	//}
	//
	//params := entityParams.(*entities.Params)

	news, err := h.service.GetAllNews(ctx)
	if err != nil {
		NewErrorResponse(ctx, http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	if len(news) == 0 {
		NewErrorResponse(ctx, http.StatusNotFound, constants.DataNotFound)
		return
	}

	NewSuccessResponse(ctx, http.StatusOK, "news search successfully", news)
	return
}

func (h *Handlers) GetNewsById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		logger.Error(err.Error(), constants.LoggerHandlers)
		NewErrorResponse(ctx, http.StatusBadRequest, constants.IdParamInvalid)
		return
	}
	
	news, err := h.service.GetNewsById(ctx, id)
	if err != nil {
		if errors.Is(err, constants.NoNewsFoundError) {
			NewErrorResponse(ctx, http.StatusNotFound, constants.NoNewsFoundError.Error())
			return
		}

		NewErrorResponse(ctx, http.StatusInternalServerError, constants.InternalServerError)
		return
	}

	NewSuccessResponse(ctx, http.StatusOK, "news search successfully", news)
}
