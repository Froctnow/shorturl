package httpserver

import (
	"github.com/gin-gonic/gin"
	metricshttp "shorturl/internal/app/httpserver/metrics"
	"shorturl/internal/app/httpserver/middleware"
	shortenhttp "shorturl/internal/app/httpserver/shorten"
	urlhttp "shorturl/internal/app/httpserver/url"
	"shorturl/internal/app/log"
	"shorturl/internal/app/usecase/metrics"
	urlusecase "shorturl/internal/app/usecase/url"
	"shorturl/internal/app/validator"
)

type ShortenerServer interface {
}

type shortenerServer struct {
	urlhttp.URLRouter
	shortenhttp.ShortenRouter
	metricshttp.MetricsRouter
}

func NewShortenerServer(
	ginEngine *gin.Engine,
	urlUseCase urlusecase.UseCase,
	logger log.LogClient,
	validator validator.Validator,
	metricsUseCase metrics.UseCase,
) ShortenerServer {
	ginEngine.Use(gin.Recovery())

	apiGroup := ginEngine.Group("/")
	apiGroup.Use(middleware.LoggingMiddleware(logger))
	apiGroup.Use(middleware.DecompressMiddleware(logger))
	apiGroup.Use(middleware.CompressMiddleware())

	return &shortenerServer{
		urlhttp.NewURLRouter(apiGroup, urlUseCase),
		shortenhttp.NewShortenRouter(apiGroup, urlUseCase, validator),
		metricshttp.NewMetricsRouter(apiGroup, metricsUseCase),
	}
}
