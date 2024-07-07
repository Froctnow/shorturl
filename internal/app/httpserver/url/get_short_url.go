package url

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/pkg/errors"

	"shorturl/internal/app/httpserver/constants"
	httpmodels "shorturl/internal/app/httpserver/models"
	usecaseerrors "shorturl/internal/app/usecase/url/errors"
)

func (r *urlRouter) GetShortURL(ctx *gin.Context) {
	alias := ctx.Param("alias")
	_, err := uuid.Parse(alias)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, httpmodels.ErrorResponse{Error: constants.MessageErrorIncorrectAlias})
		return
	}

	url, err := r.urlUseCase.GetShortURL(ctx, alias)

	if err != nil {
		if errors.As(err, &usecaseerrors.URLIsDeletedError{}) {
			ctx.Writer.WriteHeader(http.StatusGone)
			return
		} else if errors.As(err, &usecaseerrors.URLNotFound{}) {
			ctx.AbortWithStatusJSON(http.StatusNotFound, httpmodels.ErrorResponse{Error: constants.MessageErrorShortURLNotFound})
			return
		}

		ctx.AbortWithStatusJSON(http.StatusInternalServerError, httpmodels.ErrorResponse{Error: "Internal Server Error"})
		return
	}

	ctx.Redirect(http.StatusTemporaryRedirect, url)
}
