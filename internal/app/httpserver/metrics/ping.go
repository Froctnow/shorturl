package metrics

import (
	"net/http"

	"github.com/gin-gonic/gin"

	httpmodels "shorturl/internal/app/httpserver/models"
)

func (r *metricsRouter) Ping(ctx *gin.Context) {
	err := r.metricsUseCase.HealthCheckDatabaseConnection()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, httpmodels.ErrorResponse{Error: err.Error()})
		return
	}

	ctx.String(http.StatusOK, "OK")
}
