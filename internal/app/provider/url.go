package provider

import (
	"shorturl/internal/app/storage"
)

func (p *StorageProvider) CreateURL(
	url string,
) (*storage.URLEntity, error) {
	entity, err := p.storageInstance.URLRepository.CreateEntity(&storage.URLEntityDto{URL: url})

	if err != nil {
		return nil, err
	}

	return entity, nil
}

func (p *StorageProvider) GetURL(
	url string,
) *storage.URLEntity {
	entity := p.storageInstance.URLRepository.GetEntity(url)

	return entity
}
