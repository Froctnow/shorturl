package shortenerprovider

import (
	"context"
	"fmt"

	"shorturl/internal/app/shortenerprovider/models"
	"shorturl/pkg/pgclient"
)

func (p *ShortenerDBProvider) GetUserURLs(
	ctx context.Context,
	tx pgclient.Transaction,
	UserID string,
) ([]models.URL, error) {
	rows, err := p.conn.NamedQueryxContext(
		ctx,
		"GetUserURLs",
		nil,
		tx,
		UserID,
	)
	if err != nil {
		return nil, fmt.Errorf("can't execute GetUserURLS: %w", err)
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("can't execute GetUserURLS: %w", err)
	}

	return pgclient.ListValuesFromRows[models.URL](rows)
}
