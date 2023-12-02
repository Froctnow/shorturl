package services

import (
	"errors"
	"github.com/google/uuid"
)

// ServerURL TODO: Change when create config
const ServerURL = "http://localhost:8080"

type URLService struct {
	storage map[string]string
}

func NewURLService() *URLService {
	urlService := &URLService{make(map[string]string)}

	return urlService
}

func (us *URLService) CreateShortURL(url string) string {
	urlAlias := uuid.New().String()

	us.storage[urlAlias] = url

	return ServerURL + "/" + urlAlias
}

func (us *URLService) GetURL(alias string) (string, error) {
	url := us.storage[alias]

	if url == "" {
		return "", errors.New("alias not found")
	}

	return url, nil
}
