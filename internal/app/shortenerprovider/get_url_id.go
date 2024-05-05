package shortenerprovider

import (
	"context"
	"fmt"
	"shorturl/internal/app/shortenerprovider/models"
	"shorturl/pkg/pgclient"
)

func (p *ShortenerDBProvider) GetURLID(
	ctx context.Context,
	tx pgclient.Transaction,
	URL string,
) (models.URLID, error) {
	rows, err := p.conn.NamedQueryxContext(
		ctx,
		"GetURLID",
		nil,
		tx,
		URL,
	)
	if err != nil {
		return models.URLID{}, fmt.Errorf("can't execute GetURLID: %w", err)
	}

	return pgclient.StructValueFromRows[models.URLID](rows)
}
