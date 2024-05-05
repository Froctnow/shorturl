package repository

import "context"

type URL interface {
	CreateEntity(ctx context.Context, dto *URLEntityDto) (*URLEntity, error)
	GetEntity(ctx context.Context, alias string) *URLEntity
	CreateBatch(ctx context.Context, dto *[]BatchURLDto) (*[]BatchURL, error)
}

type URLEntity struct {
	ID  string
	URL string
}

type URLEntityDto struct {
	URL string
}

type BatchURL struct {
	CorrelationID string
	ShortURL      string
}

type BatchURLDto struct {
	CorrelationID string
	OriginalURL   string
}
