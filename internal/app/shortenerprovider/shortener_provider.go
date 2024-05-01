package shortenerprovider

import (
	"fmt"
	"reflect"
	"shorturl/pkg/logger"

	"shorturl/pkg/pgclient"
)

type ShortenerDBProvider struct {
	conn pgclient.PGClient
}

func NewShortenerProvider(dbConn pgclient.PGClient) ShortenerProvider {
	return &ShortenerDBProvider{
		conn: dbConn,
	}
}

func (p *ShortenerDBProvider) BeginTransaction() (pgclient.Transaction, error) {
	return p.conn.BeginTransaction()
}

func (p *ShortenerDBProvider) RollbackTransaction(tx pgclient.Transaction, log logger.LogClient) {
	if tx == nil || reflect.ValueOf(tx).IsNil() {
		return
	}
	txErr := tx.Rollback()
	if txErr != nil {
		log.Error(txErr)
	}
}

func (p *ShortenerDBProvider) CommitTransaction(tx pgclient.Transaction) error {
	if tx == nil || reflect.ValueOf(tx).IsNil() {
		return fmt.Errorf("nil transaction pointer in CommitTransaction")
	}
	return tx.Commit()
}

func (p *ShortenerDBProvider) HealthCheckConnection() error {
	return p.conn.HealthCheckConnection()
}
