package url

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"regexp"
	"shorturl/internal/app/httpserver/constants"
	httpmodels "shorturl/internal/app/httpserver/models"
	"strings"
)

func (r *urlRouter) CreateShortURL(ctx *gin.Context) {
	headerContentType := ctx.GetHeader("Content-Type")
	isCorrectHeaderContentType := r.checkHeaderContentType(headerContentType)

	if isCorrectHeaderContentType {
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

	shortURL := r.urlUseCase.CreateShortURL(url)

	ctx.String(http.StatusCreated, shortURL)
}

func (r *urlRouter) checkHeaderContentType(value string) bool {
	isTextPlain := !strings.Contains(value, "text/plain")
	isTextHtml := !strings.Contains(value, "text/html")

	return isTextPlain && isTextHtml
}
