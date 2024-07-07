package pgstorage

import (
	"context"
	"fmt"

	"shorturl/internal/app/repository"
	"shorturl/internal/app/shortenerprovider"
)

type URLRepository struct {
	provider shortenerprovider.ShortenerProvider
}

func NewURLRepository(provider shortenerprovider.ShortenerProvider) *URLRepository {
	urlRepository := &URLRepository{provider}

	return urlRepository
}

func (ur *URLRepository) CreateEntity(ctx context.Context, urlEntityDto *repository.URLEntityDto) (*repository.URLEntity, error) {
	entity, err := ur.provider.CreateURL(ctx, nil, urlEntityDto.URL, urlEntityDto.UserID)

	if err != nil {
		return nil, fmt.Errorf("can't create entity: %w", err)
	}

	if entity.ID == "" {
		r, err := ur.provider.GetURLID(ctx, nil, urlEntityDto.URL)

		if err != nil {
			return nil, fmt.Errorf("can't get entity ID: %w", err)
		}

		return nil, repository.URLDuplicateError{URL: urlEntityDto.URL, ID: r.ID}
	}

	result := repository.URLEntity(entity)

	return &result, nil
}

func (ur *URLRepository) GetEntity(ctx context.Context, alias string) *repository.URLEntity {
	entity, err := ur.provider.GetURL(ctx, nil, alias)

	if err != nil {
		return nil
	}

	result := repository.URLEntity(entity)

	return &result
}

func (ur *URLRepository) CreateBatch(ctx context.Context, batchDto *[]repository.BatchURLDto, userID string) (*[]repository.BatchURL, error) {
	entities := make([]repository.BatchURL, 0)
	tx, err := ur.provider.BeginTransaction()

	if err != nil {
		return nil, fmt.Errorf("can't begin transaction: %w", err)
	}

	for _, e := range *batchDto {
		entity, err := ur.provider.CreateURL(ctx, tx, e.OriginalURL, userID)

		if err != nil {
			err = tx.Rollback()

			if err != nil {
				return nil, fmt.Errorf("can't rollback transaction: %w", err)
			}

			return nil, fmt.Errorf("can't create entity: %w", err)
		}

		entities = append(entities, repository.BatchURL{CorrelationID: e.CorrelationID, ShortURL: entity.ID})
	}

	err = tx.Commit()

	if err != nil {
		return nil, fmt.Errorf("can't commit transaction: %w", err)
	}

	return &entities, nil
}

func (ur *URLRepository) GetUserURLs(ctx context.Context, userID string) (*[]repository.UserURL, error) {
	entities := make([]repository.UserURL, 0)
	urls, err := ur.provider.GetUserURLs(ctx, nil, userID)

	if err != nil {
		return nil, fmt.Errorf("can't get user URLs: %w", err)
	}

	for _, e := range urls {
		entities = append(entities, repository.UserURL{ShortURL: e.ID, OriginURL: e.URL})
	}

	return &entities, nil
}

func (ur *URLRepository) DeleteShortURLs(ctx context.Context, urls *[]string, userID string) error {
	err := ur.provider.DeleteURLs(ctx, nil, urls, userID)

	if err != nil {
		return fmt.Errorf("can't delete URLs: %w", err)
	}

	return nil
}
