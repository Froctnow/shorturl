package httpserver

import (
	"shorturl/internal/app/httpserver/middleware"
	shortenhttp "shorturl/internal/app/httpserver/shorten"
	urlhttp "shorturl/internal/app/httpserver/url"
	urlusecase "shorturl/internal/app/usecase/url"
	"shorturl/internal/app/validator"
	"shorturl/pkg/logger"

	"github.com/gin-gonic/gin"
)

type ShortenerServer interface {
}

type shortenerServer struct {
	urlRouter       urlhttp.Router
	shortenerRouter shortenhttp.Router
}

func NewShortenerServer(
	ginEngine *gin.Engine,
	urlUseCase urlusecase.UseCase,
	logger logger.LogClient,
	validator validator.Validator,
) ShortenerServer {
	ginEngine.Use(gin.Recovery())

	apiGroup := ginEngine.Group("/")
	apiGroup.Use(middleware.LoggingMiddleware(logger))
	apiGroup.Use(middleware.DecompressMiddleware(logger))
	apiGroup.Use(middleware.CompressMiddleware())

	return &shortenerServer{
		urlhttp.NewRouter(apiGroup, urlUseCase),
		shortenhttp.NewRouter(apiGroup, urlUseCase, validator),
	}
}
