package url

import (
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"regexp"
	httpmodels "shorturl/internal/app/httpserver/models"
	"strings"
)

const RegexpURL = "https?:\\/\\/(www\\.)?[-a-zA-Z0-9@:%._\\+~#=]{1,256}\\.[a-zA-Z0-9()]{1,6}\\b([-a-zA-Z0-9()@:%_\\+.~#?&//=]*)"

func (r *urlRouter) CreateShortURL(ctx *gin.Context) {
	headerContentType := ctx.GetHeader("Content-Type")

	if !strings.Contains(headerContentType, "text/plain") {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, httpmodels.ErrorResponse{Error: MessageErrorIncorrectContentType})
		return
	}

	body, err := io.ReadAll(ctx.Request.Body)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, httpmodels.ErrorResponse{Error: "Something went wrong"})
		return
	}

	url := string(body)

	isMatched, err := regexp.MatchString(RegexpURL, url)

	if !isMatched {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, httpmodels.ErrorResponse{Error: MessageErrorIncorrectURL})
		return
	}

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, httpmodels.ErrorResponse{Error: "Something went wrong"})
		return
	}

	shortURL := r.urlUseCase.CreateShortURL(url)

	ctx.String(http.StatusCreated, shortURL)
}
