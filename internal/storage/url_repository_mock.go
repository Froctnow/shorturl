package storage

import (
	"github.com/google/uuid"
)

type URLRepositoryMock struct {
	table map[string]*URLEntity
}

func NewURLRepositoryMock() *URLRepositoryMock {
	urlRepositoryMock := &URLRepositoryMock{make(map[string]*URLEntity)}

	return urlRepositoryMock
}

func (ur *URLRepositoryMock) CreateEntity(urlEntityDto *URLEntityDto) *URLEntity {
	ID := uuid.New().String()

	entity := &URLEntity{ID: ID, URL: urlEntityDto.URL}

	ur.table[ID] = entity

	return entity
}

func (ur *URLRepositoryMock) GetEntity(key string) *URLEntity {
	entity := ur.table[key]

	if entity == nil {
		return nil
	}

	return *(&entity)
}
