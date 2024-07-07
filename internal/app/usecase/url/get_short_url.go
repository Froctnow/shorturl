package url

import (
	"errors"

	"golang.org/x/net/context"
)

func (u *urlUseCase) GetShortURL(ctx context.Context, alias string) (string, error) {
	urlEntity := u.urlRepository.GetEntity(ctx, alias)

	if urlEntity == nil {
		return "", errors.New("alias not found")
	}

	return urlEntity.URL, nil
}
