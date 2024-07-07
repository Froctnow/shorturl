package shorten

import (
	"net/http"
	"shorturl/internal/app/httpserver/constants"
	httpmodels "shorturl/internal/app/httpserver/models"
	"shorturl/internal/app/repository"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
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

	shortURL, err := r.urlUseCase.CreateShortURL(ctx, req.URL, ctx.GetString("user_id"))

	if err != nil && errors.As(err, &repository.URLDuplicateError{}) {
		ctx.AbortWithStatusJSON(http.StatusConflict, httpmodels.CreateURLResponse{
			Result: err.(repository.URLDuplicateError).URL,
		})
		return
	}

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, httpmodels.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, httpmodels.CreateURLResponse{
		Result: shortURL,
	})
}
