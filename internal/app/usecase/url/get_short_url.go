package url

import (
	"golang.org/x/net/context"

	usecaseerrors "shorturl/internal/app/usecase/url/errors"
)

func (u *urlUseCase) GetShortURL(ctx context.Context, alias string) (string, error) {
	urlEntity := u.urlRepository.GetEntity(ctx, alias)

	if urlEntity == nil {
		return "", usecaseerrors.URLNotFoundError{}
	}

	if urlEntity.IsDeleted {
		return "", usecaseerrors.URLIsDeletedError{}
	}

	return urlEntity.URL, nil
}
