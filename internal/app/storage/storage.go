package storage

import (
	"shorturl/internal/app/client/pg"
	"shorturl/internal/app/config"
	"shorturl/internal/app/memorystorage"
	"shorturl/internal/app/migration"
	"shorturl/internal/app/pgstorage"
	"shorturl/internal/app/repository"
	"shorturl/internal/app/shortenerprovider"
	"shorturl/pkg/logger"
)

type Storage struct {
	URLRepository repository.URL
}

func NewStorage(storageMode string, cfg *config.Values, logger logger.LogClient) (*Storage, shortenerprovider.ShortenerProvider) {
	switch storageMode {
	case config.StorageModeDatabase:
		return initPgStorage(cfg, logger)
	case config.StorageModeMemory:
		return initMemoryStorage(cfg, logger), nil
	}

	return nil, nil
}

func initMemoryStorage(cfg *config.Values, logger logger.LogClient) *Storage {
	memoryStorage := memorystorage.NewStorage(cfg.FileStoragePath, logger)
	instanceStorage := &Storage{URLRepository: memoryStorage.URLRepository}

	return instanceStorage
}

func initPgStorage(cfg *config.Values, logger logger.LogClient) (*Storage, shortenerprovider.ShortenerProvider) {
	err := migration.ExecuteMigrations(cfg, logger)

	if err != nil {
		logger.Fatal(err)
	}

	shortenerDBConn, err := pg.New(cfg, logger)
	if err != nil {
		logger.Fatal(err)
	}

	shortenerProvider := shortenerprovider.NewShortenerProvider(shortenerDBConn)
	pgStorage := pgstorage.NewStorage(shortenerProvider)
	instanceStorage := &Storage{URLRepository: pgStorage.URLRepository}

	return instanceStorage, shortenerProvider
}
