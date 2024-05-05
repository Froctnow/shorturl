package repository

import (
	"context"
	"fmt"
)

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

type URLDuplicateError struct {
	URL string
	ID  string
}

func (e URLDuplicateError) Error() string {
	return fmt.Sprintf("duplicate URL - %s", e.URL)
}
