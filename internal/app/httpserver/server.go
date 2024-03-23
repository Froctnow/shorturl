package httpserver

import (
	"github.com/gin-gonic/gin"
	"shorturl/internal/app/httpserver/middleware"
	urlhttp "shorturl/internal/app/httpserver/url"
	"shorturl/internal/app/log"
	urlusecase "shorturl/internal/app/usecase/url"
)

type ShortenerServer interface {
}

type shortenerServer struct {
	urlhttp.URLRouter
}

func NewShortenerServer(
	ginEngine *gin.Engine,
	urlUseCase urlusecase.UseCase,
	logger log.LogClient,
) ShortenerServer {

	ginEngine.Use(gin.Recovery())

	apiGroup := ginEngine.Group("/")
	apiGroup.Use(middleware.LoggingMiddleware(logger))

	return &shortenerServer{
		urlhttp.NewURLRouter(apiGroup, urlUseCase),
	}
}
