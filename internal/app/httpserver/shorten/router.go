package shorten

import (
	"shorturl/internal/app/usecase/url"
	"shorturl/internal/app/validator"
	"strings"

	"github.com/gin-gonic/gin"
)

type Router interface {
	CreateShortURL(c *gin.Context)
}

type shortenRouter struct {
	urlUseCase url.UseCase
	validator  validator.Validator
}

func NewRouter(
	ginGroup *gin.RouterGroup,
	urlUseCase url.UseCase,
	validator validator.Validator,
) Router {
	router := &shortenRouter{
		urlUseCase: urlUseCase,
		validator:  validator,
	}

	urlGroup := ginGroup.Group("/api")
	urlGroup.POST("/shorten", router.CreateShortURL)
	urlGroup.POST("/shorten/batch", router.CreateBatchShortURL)

	return router
}

func (r *shortenRouter) checkHeaderContentType(value string) bool {
	isTextPlain := strings.Contains(value, "application/json")
	isXGzip := strings.Contains(value, "application/x-gzip")

	return isTextPlain || isXGzip
}
