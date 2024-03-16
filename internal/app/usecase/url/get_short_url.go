package url

import "errors"

func (u *urlUseCase) GetShortURL(alias string) (string, error) {
	urlEntity := u.provider.GetURL(alias)

	if urlEntity == nil {
		return "", errors.New("alias not found")
	}

	return urlEntity.URL, nil
}
