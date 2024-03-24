package shorten

import (
	"github.com/gin-gonic/gin"
	"shorturl/internal/app/usecase/url"
	"shorturl/internal/app/validator"
)

type ShortenRouter interface {
	CreateShortURL(c *gin.Context)
}

type shortenRouter struct {
	urlUseCase url.UseCase
	validator  validator.Validator
}

func NewShortenRouter(
	ginGroup *gin.RouterGroup,
	urlUseCase url.UseCase,
	validator validator.Validator,
) ShortenRouter {
	router := &shortenRouter{
		urlUseCase: urlUseCase,
		validator:  validator,
	}

	urlGroup := ginGroup.Group("/api")
	urlGroup.POST("/shorten", router.CreateShortURL)

	return router
}
