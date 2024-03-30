package shorten

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"shorturl/internal/app/httpserver/constants"
	httpmodels "shorturl/internal/app/httpserver/models"
	"slices"
)

func (r *shortenRouter) CreateShortURL(ctx *gin.Context) {
	contentTypeHeadersAllowed := []string{"application/x-gzip", "application/json", "text/html"}
	headerContentType := ctx.GetHeader("Content-Type")

	if !slices.Contains(contentTypeHeadersAllowed, headerContentType) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, httpmodels.ErrorResponse{Error: constants.MessageErrorIncorrectContentType})
		return
	}

	var req httpmodels.CreateURLRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, httpmodels.ErrorResponse{Error: err.Error()})
		return
	}

	errs := r.validator.ShortenCreateShortURL(&req, constants.RegexpURL)
	if len(errs.Errors) != 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, httpmodels.ErrorResponse{Error: errs.Error()})
		return
	}

	shortURL := r.urlUseCase.CreateShortURL(req.URL)

	ctx.JSON(http.StatusCreated, httpmodels.CreateURLResponse{
		Result: shortURL,
	})
}
