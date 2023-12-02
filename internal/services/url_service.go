package services

import (
	"errors"
	"github.com/google/uuid"
)

// ServerUrl TODO: Change when create config
const ServerUrl = "http://localhost:3333"

type UrlService struct {
	storage map[string]string
}

func NewUrlService() *UrlService {
	urlService := &UrlService{make(map[string]string)}

	return urlService
}

func (us *UrlService) CreateShortUrl(url string) string {
	urlAlias := uuid.New().String()

	us.storage[urlAlias] = url

	return ServerUrl + "/" + urlAlias
}

func (us *UrlService) GetUrl(alias string) (string, error) {
	url := us.storage[alias]

	if url == "" {
		return "", errors.New("alias not found")
	}

	return url, nil
}
