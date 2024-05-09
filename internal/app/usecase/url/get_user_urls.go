package url

import (
	"context"
	"fmt"
	"shorturl/internal/app/repository"
)

func (u *urlUseCase) GetUserURLs(ctx context.Context, userID string) (*[]repository.UserURL, error) {
	urls, err := u.urlRepository.GetUserURLs(ctx, userID)

	if err != nil {
		u.logger.ErrorCtx(ctx, fmt.Errorf("can't get user urls: %w", err), "user_id", userID)
		return nil, err
	}

	result := make([]repository.UserURL, 0)

	for _, r := range *urls {
		result = append(result, repository.UserURL{
			OriginURL: r.OriginURL,
			ShortURL:  u.serverURL + "/" + r.ShortURL,
		})
	}

	return &result, nil
}
