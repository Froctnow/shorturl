package repository

import (
	"context"
	"fmt"
)

type URL interface {
	CreateEntity(ctx context.Context, dto *URLEntityDto) (*URLEntity, error)
	GetEntity(ctx context.Context, alias string) *URLEntity
	CreateBatch(ctx context.Context, dto *[]BatchURLDto, userId string) (*[]BatchURL, error)
	GetUserURLs(ctx context.Context, userID string) (*[]UserURL, error)
	DeleteShortURLs(ctx context.Context, urls *[]string, userID string) error
}

type URLEntity struct {
	ID        string
	URL       string
	UserID    string
	IsDeleted bool
}

type URLEntityDto struct {
	URL    string
	UserID string
}

type BatchURL struct {
	CorrelationID string
	ShortURL      string
}

type BatchURLDto struct {
	CorrelationID string
	OriginalURL   string
}

type UserURL struct {
	ShortURL  string
	OriginURL string
}

type URLDuplicateError struct {
	URL string
	ID  string
}

func (e URLDuplicateError) Error() string {
	return fmt.Sprintf("duplicate URL - %s", e.URL)
}
