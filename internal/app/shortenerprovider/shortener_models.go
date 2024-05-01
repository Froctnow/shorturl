package shortenerprovider

import (
	"context"
	"shorturl/internal/app/shortenerprovider/models"
	"shorturl/pkg/logger"
	"shorturl/pkg/pgclient"
)

type ShortenerProvider interface {
	CreateURL(
		ctx context.Context,
		tx pgclient.Transaction,
		URL string,
	) (models.URL, error)

	GetURL(
		ctx context.Context,
		tx pgclient.Transaction,
		alias string,
	) (models.URL, error)

	BeginTransaction() (pgclient.Transaction, error)
	RollbackTransaction(tx pgclient.Transaction, log logger.LogClient)
	CommitTransaction(tx pgclient.Transaction) error
	HealthCheckConnection() error
}
