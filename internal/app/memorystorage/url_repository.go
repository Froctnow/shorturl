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

	entity := &repository.URLEntity{ID: ID, URL: urlEntityDto.URL, UserID: urlEntityDto.UserID}

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

func (ur *URLRepository) CreateBatch(_ context.Context, batchDto *[]repository.BatchURLDto, userID string) (*[]repository.BatchURL, error) {
	result := make([]repository.BatchURL, 0)

	for _, urlDto := range *batchDto {
		ID := uuid.New().String()

		entity := &repository.URLEntity{ID: ID, URL: urlDto.OriginalURL, UserID: userID}

		err := ur.writeToFile(entity)
		if err != nil {
			return nil, fmt.Errorf("can't save entity, err %w", err)
		}

		ur.table[ID] = entity

		result = append(result, repository.BatchURL{CorrelationID: urlDto.CorrelationID, ShortURL: ID})
	}

	return &result, nil
}

func (ur *URLRepository) GetUserURLs(_ context.Context, userID string) (*[]repository.UserURL, error) {
	result := make([]repository.UserURL, 0)

	for _, entity := range ur.table {
		if entity.UserID == userID {
			result = append(result, repository.UserURL{ShortURL: entity.ID, OriginURL: entity.URL})
		}
	}

	return &result, nil
}

func (ur *URLRepository) DeleteShortURLs(_ context.Context, request []string, userID string) error {
	for _, shortURL := range request {
		entity := ur.table[shortURL]

		if entity == nil {
			fmt.Printf("can't find entity with short URL %s \n", shortURL)
			continue
		}

		if entity.UserID != userID {
			fmt.Printf("entity with short URL %s doesn't belong to user %s \n", shortURL, userID)
			continue
		}

		entity.IsDeleted = true
	}

	return nil
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

	URLFromFileJSON, err := json.Marshal(models.URLFromFile{UUID: urlEntity.ID, ShortURL: urlEntity.ID, OriginURL: urlEntity.URL, UserID: urlEntity.UserID})
	if err != nil {
		return fmt.Errorf("can't marshal URL, err %w", err)
	}

	_, err = file.WriteString(string(URLFromFileJSON) + "\n")
	if err != nil {
		return fmt.Errorf("can't write to file storage, err %w", err)
	}

	return nil
}
