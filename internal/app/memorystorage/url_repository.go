package memorystorage

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"shorturl/internal/app/memorystorage/models"
	"shorturl/internal/app/repository"

	"github.com/google/uuid"
)

type URLRepository struct {
	table           map[string]*repository.URLEntity
	storageFilePath string
}

func NewURLRepository(storageFilePath string) *URLRepository {
	urlRepository := &URLRepository{make(map[string]*repository.URLEntity), storageFilePath}

	return urlRepository
}

func (ur *URLRepository) CreateEntity(_ context.Context, urlEntityDto *repository.URLEntityDto) (*repository.URLEntity, error) {
	ID := uuid.New().String()

	entity := &repository.URLEntity{ID: ID, URL: urlEntityDto.URL}

	err := ur.writeToFile(entity)

	if err != nil {
		return nil, fmt.Errorf("can't save entity, err %w", err)
	}

	ur.table[ID] = entity

	return entity, nil
}

func (ur *URLRepository) GetEntity(_ context.Context, key string) *repository.URLEntity {
	entity := ur.table[key]

	return entity
}

func (ur *URLRepository) AddEntity(urlEntity *repository.URLEntity) {
	ur.table[urlEntity.ID] = urlEntity
}

func (ur *URLRepository) writeToFile(urlEntity *repository.URLEntity) error {
	if ur.storageFilePath == "" {
		return nil
	}

	file, err := os.OpenFile(ur.storageFilePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)

	if err != nil {
		return fmt.Errorf("can't open file of storage, err %w", err)
	}

	defer file.Close()

	URLFromFileJSON, err := json.Marshal(models.URLFromFile{UUID: urlEntity.ID, ShortURL: urlEntity.ID, OriginURL: urlEntity.URL})

	if err != nil {
		return fmt.Errorf("can't marshal URL, err %w", err)
	}

	_, err = file.WriteString(string(URLFromFileJSON) + "\n")

	if err != nil {
		return fmt.Errorf("can't write to file storage, err %w", err)
	}

	return nil
}
