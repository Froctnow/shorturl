package storage

import (
	"github.com/google/uuid"
)

type IURLRepository interface {
	CreateEntity(*URLEntityDto) *URLEntity
	GetEntity(string) *URLEntity
}

type URLEntity struct {
	ID  string
	URL string
}

type URLEntityDto struct {
	URL string
}

type URLRepository struct {
	table map[string]*URLEntity
}

func NewURLRepository() *URLRepository {
	urlRepository := &URLRepository{make(map[string]*URLEntity)}

	return urlRepository
}

func (ur *URLRepository) CreateEntity(urlEntityDto *URLEntityDto) *URLEntity {
	ID := uuid.New().String()

	entity := &URLEntity{ID: ID, URL: urlEntityDto.URL}

	ur.table[ID] = entity

	return entity
}

func (ur *URLRepository) GetEntity(key string) *URLEntity {
	entity := ur.table[key]

	return entity
}
