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
		UserID string,
	) (models.URL, error)

	GetURL(
		ctx context.Context,
		tx pgclient.Transaction,
		alias string,
	) (models.URL, error)

	GetURLID(
		ctx context.Context,
		tx pgclient.Transaction,
		alias string,
	) (models.URLID, error)

	GetUserURLs(
		ctx context.Context,
		tx pgclient.Transaction,
		UserID string,
	) ([]models.URL, error)

	DeleteURLs(
		ctx context.Context,
		tx pgclient.Transaction,
		urls []string,
		userID string,
	) error

	BeginTransaction() (pgclient.Transaction, error)
	RollbackTransaction(tx pgclient.Transaction, log logger.LogClient)
	CommitTransaction(tx pgclient.Transaction) error
	HealthCheckConnection() error
}
