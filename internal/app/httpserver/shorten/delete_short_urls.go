package shorten

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"shorturl/internal/app/httpserver/constants"
	httpmodels "shorturl/internal/app/httpserver/models"
)

func (r *shortenRouter) DeleteShortURLs(ctx *gin.Context) {
	headerContentType := ctx.GetHeader("Content-Type")
	isCorrectHeaderContentType := r.checkHeaderContentType(headerContentType)

	if !isCorrectHeaderContentType {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, httpmodels.ErrorResponse{Error: constants.MessageErrorIncorrectContentType})
		return
	}

	var req []string
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, httpmodels.ErrorResponse{Error: err.Error()})
		return
	}

	errs := r.validator.ShortenDeleteShortURLs(&req)
	if len(errs.Errors) != 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, httpmodels.ErrorResponse{Error: errs.Error()})
		return
	}

	go func() {
		r.urlUseCase.DeleteShortURLs(ctx, req, ctx.GetString("user_id"))
	}()

	ctx.Writer.WriteHeader(http.StatusAccepted)
}
