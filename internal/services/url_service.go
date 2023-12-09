package services

import (
	"errors"
	"shorturl/internal/storage"
)

// ServerURL TODO: Change when create config
const ServerURL = "http://localhost:8080"

type URLService struct {
	urlRepository storage.IURLRepository
}

func NewURLService(urlRepository storage.IURLRepository) *URLService {
	urlService := &URLService{urlRepository: urlRepository}

	return urlService
}

func (us *URLService) CreateShortURL(url string) string {
	urlEntityDto := storage.URLEntityDto{URL: url}
	urlEntity := us.urlRepository.CreateEntity(&urlEntityDto)

	return ServerURL + "/" + urlEntity.ID
}

func (us *URLService) GetURL(alias string) (string, error) {
	urlEntity := us.urlRepository.GetEntity(alias)

	if urlEntity == nil {
		return "", errors.New("alias not found")
	}

	return urlEntity.URL, nil
}
