package url

import (
	"net/http"
	"shorturl/internal/app/httpserver/constants"
	httpmodels "shorturl/internal/app/httpserver/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (r *urlRouter) GetShortURL(ctx *gin.Context) {
	alias := ctx.Param("alias")
	_, err := uuid.Parse(alias)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, httpmodels.ErrorResponse{Error: constants.MessageErrorIncorrectAlias})
		return
	}

	url, err := r.urlUseCase.GetShortURL(alias)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, httpmodels.ErrorResponse{Error: constants.MessageErrorShortURLNotFound})
		return
	}

	ctx.Redirect(http.StatusTemporaryRedirect, url)
}
