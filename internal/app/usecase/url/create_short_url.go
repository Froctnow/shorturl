package url

import (
	"context"
	"shorturl/internal/app/repository"
)

func (u *urlUseCase) CreateShortURL(ctx context.Context, url string) (string, error) {
	urlEntity, err := u.urlRepository.CreateEntity(ctx, &repository.URLEntityDto{URL: url})

	if err != nil {
		return "", err
	}

	return u.serverURL + "/" + urlEntity.ID, nil
}
