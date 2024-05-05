package url

import (
	httpmodels "shorturl/internal/app/httpserver/models"
	"shorturl/internal/app/repository"

	"golang.org/x/net/context"
)

type urlUseCase struct {
	urlRepository repository.URL
	serverURL     string
}

func NewUseCase(
	urlRepository repository.URL,
	serverURL string,
) UseCase {
	return &urlUseCase{
		urlRepository: urlRepository,
		serverURL:     serverURL,
	}
}

type UseCase interface {
	CreateShortURL(ctx context.Context, url string) (string, error)
	GetShortURL(ctx context.Context, alias string) (string, error)
	CreateBatchShortURL(ctx context.Context, request *[]httpmodels.CreateBatchURLRequest) (*[]repository.BatchURL, error)
}
