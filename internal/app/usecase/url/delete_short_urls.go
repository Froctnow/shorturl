package url

import (
	"context"
	"fmt"
)

func (u *urlUseCase) DeleteShortURLs(ctx context.Context, request []string, userID string) error {
	err := u.urlRepository.DeleteShortURLs(ctx, request, userID)
	if err != nil {
		u.logger.ErrorCtx(ctx, fmt.Errorf("failed to delete short urls: %w", err))
		return err
	}

	return nil
}
