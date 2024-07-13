package shorten

import (
	"net/http"

	"shorturl/internal/app/httpserver/constants"
	httpmodels "shorturl/internal/app/httpserver/models"

	"github.com/gin-gonic/gin"
)

func (r *shortenRouter) CreateBatchShortURL(ctx *gin.Context) {
	headerContentType := ctx.GetHeader("Content-Type")
	isCorrectHeaderContentType := r.checkHeaderContentType(headerContentType)

	if !isCorrectHeaderContentType {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, httpmodels.ErrorResponse{Error: constants.MessageErrorIncorrectContentType})
		return
	}

	var req []httpmodels.CreateBatchURLRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, httpmodels.ErrorResponse{Error: err.Error()})
		return
	}

	errs := r.validator.ShortenCreateBatchShortURL(&req)
	if len(errs.Errors) != 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, httpmodels.ErrorResponse{Error: errs.Error()})
		return
	}

	batchURL, err := r.urlUseCase.CreateBatchShortURL(ctx, &req, ctx.GetString("user_id"))
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, httpmodels.ErrorResponse{Error: err.Error()})
		return
	}

	result := make([]httpmodels.CreateBatchURLResponse, 0)

	for _, url := range *batchURL {
		result = append(result, httpmodels.CreateBatchURLResponse{
			CorrelationID: url.CorrelationID,
			ShortURL:      url.ShortURL,
		})
	}

	ctx.JSON(http.StatusCreated, &result)
}
