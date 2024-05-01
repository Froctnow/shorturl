package metrics

import (
	"shorturl/internal/app/usecase/metrics"

	"github.com/gin-gonic/gin"
)

type Router interface {
	Ping(c *gin.Context)
}

type metricsRouter struct {
	metricsUseCase metrics.UseCase
}

func NewRouter(
	ginGroup *gin.RouterGroup,
	metricsUseCase metrics.UseCase,
) Router {
	router := &metricsRouter{
		metricsUseCase: metricsUseCase,
	}

	urlGroup := ginGroup.Group("/")
	urlGroup.GET("/ping", router.Ping)

	return router
}
