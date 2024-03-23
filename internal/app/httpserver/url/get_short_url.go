package url

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	httpmodels "shorturl/internal/app/httpserver/models"
)

func (r *urlRouter) GetShortURL(ctx *gin.Context) {
	alias := ctx.Param("alias")
	_, err := uuid.Parse(alias)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, httpmodels.ErrorResponse{Error: MessageErrorIncorrectAlias})
		return
	}

	url, err := r.urlUseCase.GetShortURL(alias)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, httpmodels.ErrorResponse{Error: MessageErrorShortURLNotFound})
		return
	}

	ctx.Redirect(http.StatusTemporaryRedirect, url)
}
