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

func NewURLRepository(provider shortenerprovider.ShortenerProvider) repository.URL {
	urlRepository := &URLRepository{provider}

	return urlRepository
}

func (ur *URLRepository) CreateEntity(ctx context.Context, urlEntityDto *repository.URLEntityDto) (*repository.URLEntity, error) {
	entity, err := ur.provider.CreateURL(ctx, nil, urlEntityDto.URL)

	if err != nil {
		return nil, fmt.Errorf("can't create entity: %w", err)
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
