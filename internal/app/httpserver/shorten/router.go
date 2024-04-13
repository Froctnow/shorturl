package shorten

import (
	"shorturl/internal/app/usecase/url"
	"shorturl/internal/app/validator"

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

	return router
}
