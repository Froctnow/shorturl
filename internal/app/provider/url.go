package provider

import (
	"shorturl/internal/app/storage"
)

func (p *StorageProvider) CreateURL(
	url string,
) *storage.URLEntity {
	entity := p.storageInstance.URLRepository.CreateEntity(&storage.URLEntityDto{URL: url})

	return entity
}

func (p *StorageProvider) GetURL(
	url string,
) *storage.URLEntity {
	entity := p.storageInstance.URLRepository.GetEntity(url)

	return entity
}
