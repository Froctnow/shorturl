package url

import (
	"github.com/gin-gonic/gin"
	"shorturl/internal/app/usecase/url"
)

type URLRouter interface {
	CreateShortURL(c *gin.Context)
	GetShortURL(c *gin.Context)
}

type urlRouter struct {
	urlUseCase url.UseCase
}

func NewURLRouter(
	ginGroup *gin.RouterGroup,
	urlUseCase url.UseCase,
) URLRouter {
	router := &urlRouter{
		urlUseCase: urlUseCase,
	}

	urlGroup := ginGroup.Group("/")
	urlGroup.POST("/", router.CreateShortURL)
	urlGroup.GET("/:alias", router.GetShortURL)

	return router
}
