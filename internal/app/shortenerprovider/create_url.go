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
	URL string,
) (models.URL, error) {
	rows, err := p.conn.NamedQueryxContext(
		ctx,
		"CreateURL",
		nil,
		tx,
		URL,
	)
	if err != nil {
		return models.URL{}, fmt.Errorf("can't execute CreateURL: %w", err)
	}
	return pgclient.StructValueFromRows[models.URL](rows)
}
