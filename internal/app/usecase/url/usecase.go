package url

import (
	"shorturl/internal/app/provider"
)

type urlUseCase struct {
	provider  provider.IStorageProvider
	serverURL string
}

func NewUseCase(
	provider provider.IStorageProvider,
	serverURL string,
) UseCase {
	return &urlUseCase{
		provider:  provider,
		serverURL: serverURL,
	}
}

type UseCase interface {
	CreateShortURL(url string) (string, error)
	GetShortURL(alias string) (string, error)
}
