package metrics

import (
	"shorturl/internal/app/provider"
)

type metricsUseCase struct {
	provider provider.ShortenerProvider
}

func NewUseCase(
	provider provider.ShortenerProvider,
) UseCase {
	return &metricsUseCase{
		provider: provider,
	}
}

type UseCase interface {
	HealthCheckDatabaseConnection() error
}
