package shortenerprovider

import (
	"context"
	"fmt"

	"shorturl/internal/app/shortenerprovider/models"
	"shorturl/pkg/pgclient"
)

func (p *ShortenerDBProvider) CreateURL(
	ctx context.Context,
	tx pgclient.Transaction,
	url string,
) (models.URL, error) {
	rows, err := p.conn.NamedQueryxContext(
		ctx,
		"CreateURL",
		nil,
		tx,
		url,
	)
	if err != nil {
		return models.URL{}, fmt.Errorf("can't execute CreateURL: %w", err)
	}

	err = rows.Err()

	if err != nil {
		return models.URL{}, fmt.Errorf("can't execute CreateURL: %w", err)
	}

	return pgclient.StructValueFromRows[models.URL](rows)
}
