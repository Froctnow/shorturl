package httpserver

import (
	"github.com/gin-gonic/gin"
	"shorturl/internal/app/httpserver/middleware"
	shortenhttp "shorturl/internal/app/httpserver/shorten"
	urlhttp "shorturl/internal/app/httpserver/url"
	"shorturl/internal/app/log"
	urlusecase "shorturl/internal/app/usecase/url"
	"shorturl/internal/app/validator"
)

type ShortenerServer interface {
}

type shortenerServer struct {
	urlhttp.URLRouter
	shortenhttp.ShortenRouter
}

func NewShortenerServer(
	ginEngine *gin.Engine,
	urlUseCase urlusecase.UseCase,
	logger log.LogClient,
	validator validator.Validator,
) ShortenerServer {
	ginEngine.Use(gin.Recovery())

	apiGroup := ginEngine.Group("/")
	apiGroup.Use(middleware.LoggingMiddleware(logger))
	apiGroup.Use(middleware.DecompressMiddleware(logger))
	apiGroup.Use(middleware.CompressMiddleware())

	return &shortenerServer{
		urlhttp.NewURLRouter(apiGroup, urlUseCase),
		shortenhttp.NewShortenRouter(apiGroup, urlUseCase, validator),
	}
}
