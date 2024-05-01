package repository

import "context"

type URL interface {
	CreateEntity(ctx context.Context, dto *URLEntityDto) (*URLEntity, error)
	GetEntity(ctx context.Context, alias string) *URLEntity
}

type URLEntity struct {
	ID  string
	URL string
}

type URLEntityDto struct {
	URL string
}
