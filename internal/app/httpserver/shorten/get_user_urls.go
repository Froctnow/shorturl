package shorten

import (
	"net/http"
	"shorturl/internal/app/httpserver/constants"
	httpmodels "shorturl/internal/app/httpserver/models"

	"github.com/gin-gonic/gin"
)

func (r *shortenRouter) GetUserURLS(ctx *gin.Context) {
	isNewUser := ctx.GetBool(constants.ContextIsNewUser)

	if isNewUser {
		ctx.Status(http.StatusUnauthorized)
		return
	}

	userID := ctx.GetString(constants.ContextUserID)
	urls, err := r.urlUseCase.GetUserURLs(ctx, userID)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, httpmodels.ErrorResponse{Error: err.Error()})
		return
	}

	if len(*urls) == 0 {
		ctx.Status(http.StatusNoContent)
		return
	}

	response := make([]httpmodels.GetUserURLsResponse, 0)

	for _, url := range *urls {
		response = append(response, httpmodels.GetUserURLsResponse{
			OriginalURL: url.OriginURL,
			ShortURL:    url.ShortURL,
		})
	}

	ctx.JSON(http.StatusOK, response)
}
