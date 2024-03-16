package url

import (
	"shorturl/internal/app/provider"
)

type urlUseCase struct {
	provider provider.IStorageProvider
}

func NewUseCase(
	provider provider.IStorageProvider,
) UseCase {
	return &urlUseCase{
		provider: provider,
	}
}

type UseCase interface {
	CreateShortURL(url string) string
	GetShortURL(alias string) (string, error)
}
