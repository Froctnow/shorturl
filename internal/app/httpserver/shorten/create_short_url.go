package shorten

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"shorturl/internal/app/httpserver/constants"
	httpmodels "shorturl/internal/app/httpserver/models"
	"strings"
)

func (r *shortenRouter) CreateShortURL(ctx *gin.Context) {
	headerContentType := ctx.GetHeader("Content-Type")

	if !strings.Contains(headerContentType, "application/json") {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, httpmodels.ErrorResponse{Error: constants.MessageErrorIncorrectContentType})
		return
	}

	var req httpmodels.CreateUrlRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		fmt.Println("ERROR THERE 1", err.Error())
		ctx.AbortWithStatusJSON(http.StatusBadRequest, httpmodels.ErrorResponse{Error: err.Error()})
		return
	}

	errs := r.validator.ShortenCreateShortURL(&req, constants.RegexpURL)
	if len(errs.Errors) != 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, httpmodels.ErrorResponse{Error: errs.Error()})
		return
	}

	shortURL := r.urlUseCase.CreateShortURL(req.URL)

	ctx.JSON(http.StatusCreated, httpmodels.CreateUrlResponse{
		Result: shortURL,
	})
}
