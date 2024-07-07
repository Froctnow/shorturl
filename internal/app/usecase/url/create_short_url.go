package url

import (
	"context"
	"fmt"
	"shorturl/internal/app/repository"

	"github.com/pkg/errors"
)

func (u *urlUseCase) CreateShortURL(ctx context.Context, url string, userId string) (string, error) {
	urlEntity, err := u.urlRepository.CreateEntity(ctx, &repository.URLEntityDto{URL: url, UserID: userId})

	if err != nil && errors.As(err, &repository.URLDuplicateError{}) {
		u.logger.InfoCtx(ctx, "URL already exists", "url", url, "user_id", userId, "error", err)
		return "", repository.URLDuplicateError{URL: u.serverURL + "/" + err.(repository.URLDuplicateError).ID}
	}

	if err != nil {
		u.logger.ErrorCtx(ctx, fmt.Errorf("can't create entity: %w", err), "url", url, "user_id", userId)
		return "", err
	}

	return u.serverURL + "/" + urlEntity.ID, nil
}
