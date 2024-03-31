package metrics

import (
	"github.com/gin-gonic/gin"
	"shorturl/internal/app/usecase/metrics"
)

type MetricsRouter interface {
	Ping(c *gin.Context)
}

type metricsRouter struct {
	metricsUseCase metrics.UseCase
}

func NewMetricsRouter(
	ginGroup *gin.RouterGroup,
	metricsUseCase metrics.UseCase,
) MetricsRouter {
	router := &metricsRouter{
		metricsUseCase: metricsUseCase,
	}

	urlGroup := ginGroup.Group("/")
	urlGroup.GET("/ping", router.Ping)

	return router
}
