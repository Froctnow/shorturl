package metrics

import (
	"shorturl/internal/app/shortenerprovider"
)

type metricsUseCase struct {
	provider shortenerprovider.ShortenerProvider
}

func NewUseCase(
	provider shortenerprovider.ShortenerProvider,
) UseCase {
	return &metricsUseCase{
		provider: provider,
	}
}

type UseCase interface {
	HealthCheckDatabaseConnection() error
}
