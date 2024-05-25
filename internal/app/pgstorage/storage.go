package pgstorage

import (
	"shorturl/internal/app/repository"
	"shorturl/internal/app/shortenerprovider"
)

type Instance struct {
	URLRepository repository.URL
}

func NewStorage(provider shortenerprovider.ShortenerProvider) *Instance {
	storage := &Instance{URLRepository: NewURLRepository(provider)}

	return storage
}
