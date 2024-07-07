package shortenerprovider

import (
	"context"
	"fmt"

	"shorturl/pkg/pgclient"
)

func (p *ShortenerDBProvider) DeleteURLs(
	ctx context.Context,
	tx pgclient.Transaction,
	urls *[]string,
	userID string,
) error {
	_, err := p.conn.NamedQueryxContext(
		ctx,
		"DeleteURLs",
		p.mapper.URLIDs(urls),
		tx,
		userID,
	)
	if err != nil {
		return fmt.Errorf("can't execute DeleteURLs: %w", err)
	}

	return nil
}
