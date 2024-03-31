package provider

import (
	"shorturl/internal/app/log"
	"shorturl/pkg/pgclient"
)

type ShortenerProvider interface {
	BeginTransaction() (pgclient.Transaction, error)
	RollbackTransaction(tx pgclient.Transaction, log log.LogClient)
	CommitTransaction(tx pgclient.Transaction) error
	HealthCheckConnection() error
}
