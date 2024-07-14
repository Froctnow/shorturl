package url

import (
	"errors"
	"io"
	"net/http"
	"regexp"
	"strings"

	"shorturl/internal/app/httpserver/constants"
	httpmodels "shorturl/internal/app/httpserver/models"
	"shorturl/internal/app/repository"

	"github.com/gin-gonic/gin"
)

func (r *urlRouter) CreateShortURL(ctx *gin.Context) {
	headerContentType := ctx.GetHeader("Content-Type")
	isCorrectHeaderContentType := r.checkHeaderContentType(headerContentType)

	if !isCorrectHeaderContentType {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, httpmodels.ErrorResponse{Error: constants.MessageErrorIncorrectContentType})
		return
	}

	body, err := io.ReadAll(ctx.Request.Body)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, httpmodels.ErrorResponse{Error: "Something went wrong"})
		return
	}

	url := string(body)

	isMatched, err := regexp.MatchString(constants.RegexpURL, url)

	if !isMatched {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, httpmodels.ErrorResponse{Error: constants.MessageErrorIncorrectURL})
		return
	}

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, httpmodels.ErrorResponse{Error: "Something went wrong"})
		return
	}

	shortURL, err := r.urlUseCase.CreateShortURL(ctx, url, ctx.GetString("user_id"))

	if err != nil && errors.As(err, &repository.URLDuplicateError{}) {
		ctx.String(http.StatusConflict, err.(repository.URLDuplicateError).URL)
		return
	}

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, httpmodels.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.String(http.StatusCreated, shortURL)
}

func (r *urlRouter) checkHeaderContentType(value string) bool {
	isTextPlain := strings.Contains(value, "text/plain")
	isTextHTML := strings.Contains(value, "text/html")
	isXGzip := strings.Contains(value, "application/x-gzip")

	return isTextPlain || isTextHTML || isXGzip
}
