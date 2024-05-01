package shorten

import (
	"net/http"
	"shorturl/internal/app/httpserver/constants"
	httpmodels "shorturl/internal/app/httpserver/models"
	"strings"

	"github.com/gin-gonic/gin"
)

func (r *shortenRouter) CreateShortURL(ctx *gin.Context) {
	headerContentType := ctx.GetHeader("Content-Type")
	isCorrectHeaderContentType := r.checkHeaderContentType(headerContentType)

	if !isCorrectHeaderContentType {
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

	shortURL, err := r.urlUseCase.CreateShortURL(ctx, req.URL)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, httpmodels.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, httpmodels.CreateURLResponse{
		Result: shortURL,
	})
}

func (r *shortenRouter) checkHeaderContentType(value string) bool {
	isTextPlain := strings.Contains(value, "application/json")
	isXGzip := strings.Contains(value, "application/x-gzip")

	return isTextPlain || isXGzip
}
