package url

import (
	httpmodels "shorturl/internal/app/httpserver/models"
	"shorturl/internal/app/repository"
	"shorturl/pkg/logger"

	"golang.org/x/net/context"
)

type urlUseCase struct {
	urlRepository repository.URL
	serverURL     string
	logger        logger.LogClient
}

func NewUseCase(
	urlRepository repository.URL,
	serverURL string,
	logger logger.LogClient,
) UseCase {
	return &urlUseCase{
		urlRepository: urlRepository,
		serverURL:     serverURL,
		logger:        logger,
	}
}

type UseCase interface {
	CreateShortURL(ctx context.Context, url string, userID string) (string, error)
	GetShortURL(ctx context.Context, alias string) (string, error)
	CreateBatchShortURL(ctx context.Context, request *[]httpmodels.CreateBatchURLRequest, userID string) (*[]repository.BatchURL, error)
	GetUserURLs(ctx context.Context, userID string) (*[]repository.UserURL, error)
}
