package url

import (
	"context"
	httpmodels "shorturl/internal/app/httpserver/models"
	"shorturl/internal/app/repository"
)

func (u *urlUseCase) CreateBatchShortURL(ctx context.Context, request *[]httpmodels.CreateBatchURLRequest, userID string) (*[]repository.BatchURL, error) {
	dto := make([]repository.BatchURLDto, 0)

	for _, r := range *request {
		dto = append(dto, repository.BatchURLDto{
			CorrelationID: r.CorrelationID,
			OriginalURL:   r.OriginalURL,
		})
	}

	batchURL, err := u.urlRepository.CreateBatch(ctx, &dto, userID)

	if err != nil {
		return nil, err
	}

	result := make([]repository.BatchURL, 0)

	for _, r := range *batchURL {
		result = append(result, repository.BatchURL{
			CorrelationID: r.CorrelationID,
			ShortURL:      u.serverURL + "/" + r.ShortURL,
		})
	}

	return &result, nil
}
