package httpserver

import (
	"shorturl/internal/app/config"
	metricshttp "shorturl/internal/app/httpserver/metrics"
	"shorturl/internal/app/httpserver/middleware"
	shortenhttp "shorturl/internal/app/httpserver/shorten"
	urlhttp "shorturl/internal/app/httpserver/url"
	"shorturl/internal/app/usecase/metrics"
	urlusecase "shorturl/internal/app/usecase/url"
	"shorturl/internal/app/validator"
	"shorturl/pkg/logger"

	"github.com/gin-gonic/gin"
)

type ShortenerServer interface{}

type shortenerServer struct {
	urlRouter       urlhttp.Router
	shortenerRouter shortenhttp.Router
	metricsRouter   metricshttp.Router
}

func NewShortenerServer(
	ginEngine *gin.Engine,
	urlUseCase urlusecase.UseCase,
	logger logger.LogClient,
	validator validator.Validator,
	metricsUseCase metrics.UseCase,
	cfg *config.Values,
) ShortenerServer {
	ginEngine.Use(gin.Recovery())

	apiGroup := ginEngine.Group("/")
	apiGroup.Use(middleware.AccessControlMiddleware(cfg, logger))
	apiGroup.Use(middleware.LoggingMiddleware(logger))
	apiGroup.Use(middleware.DecompressMiddleware(logger))
	apiGroup.Use(middleware.CompressMiddleware())

	return &shortenerServer{
		urlhttp.NewRouter(apiGroup, urlUseCase),
		shortenhttp.NewRouter(apiGroup, urlUseCase, validator),
		metricshttp.NewRouter(apiGroup, metricsUseCase),
	}
}
