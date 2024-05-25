package url

import (
	"context"
	"shorturl/internal/app/repository"

	"github.com/pkg/errors"
)

func (u *urlUseCase) CreateShortURL(ctx context.Context, url string) (string, error) {
	urlEntity, err := u.urlRepository.CreateEntity(ctx, &repository.URLEntityDto{URL: url})

	if err != nil && errors.As(err, &repository.URLDuplicateError{}) {
		return "", repository.URLDuplicateError{URL: u.serverURL + "/" + err.(repository.URLDuplicateError).ID}
	}

	if err != nil {
		return "", err
	}

	return u.serverURL + "/" + urlEntity.ID, nil
}
