package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"news-service/internal/entities"
	"news-service/internal/server/http/handlers"
	"news-service/pkg/constants"
)

func ParseQuery() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ctx.Query("id")
		description := ctx.Query("desc")
		title := ctx.Query("title")
		author := ctx.Query("author")
		footer := ctx.Query("footer")

		if id == "" && description == "" && title == "" && author == "" && footer == "" {
			handlers.NewErrorResponse(ctx, http.StatusBadRequest, constants.ParamsIsRequired)
			return
		}

		ctx.Set("params", &entities.Params{
			DiscordId:   id,
			Title:       title,
			Description: description,
			Author:      author,
			Footer:      footer,
		})

		ctx.Next()
	}
}
