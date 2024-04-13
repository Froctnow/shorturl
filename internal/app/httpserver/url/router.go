package url

import (
	"shorturl/internal/app/usecase/url"

	"github.com/gin-gonic/gin"
)

type Router interface {
	CreateShortURL(c *gin.Context)
	GetShortURL(c *gin.Context)
}

type urlRouter struct {
	urlUseCase url.UseCase
}

func NewRouter(
	ginGroup *gin.RouterGroup,
	urlUseCase url.UseCase,
) Router {
	router := &urlRouter{
		urlUseCase: urlUseCase,
	}

	urlGroup := ginGroup.Group("/")
	urlGroup.POST("/", router.CreateShortURL)
	urlGroup.GET("/:alias", router.GetShortURL)

	return router
}
