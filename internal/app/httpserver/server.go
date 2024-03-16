package httpserver

import (
	"github.com/gin-gonic/gin"
	urlhttp "shorturl/internal/app/httpserver/url"
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
) ShortenerServer {

	ginEngine.Use(gin.Recovery())

	apiGroup := ginEngine.Group("/")

	return &shortenerServer{
		urlhttp.NewURLRouter(apiGroup, urlUseCase),
	}
}
