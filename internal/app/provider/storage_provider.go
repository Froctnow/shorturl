package provider

import (
	"shorturl/internal/app/storage"
)

type StorageProvider struct {
	storageInstance *storage.Instance
}

func NewStorageProvider(storageInstance *storage.Instance) IStorageProvider {
	return &StorageProvider{storageInstance: storageInstance}
}

type IStorageProvider interface {
	CreateURL(url string) (*storage.URLEntity, error)
	GetURL(url string) *storage.URLEntity
}
