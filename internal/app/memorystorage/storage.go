package memorystorage

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"

	"shorturl/internal/app/memorystorage/models"
	"shorturl/internal/app/repository"
	"shorturl/pkg/logger"

	"github.com/pkg/errors"
)

type Instance struct {
	URLRepository *URLRepository
}

func NewStorage(filePath string, logger logger.LogClient) *Instance {
	storage := &Instance{URLRepository: NewURLRepository(filePath)}

	if filePath != "" {
		initFromFile(filePath, storage, logger)
	}

	return storage
}

func initFromFile(storageFilePath string, storage *Instance, logger logger.LogClient) {
	logger.Info("Start init storage from file")
	_, err := os.Stat(storageFilePath)

	if err != nil && errors.Is(err, os.ErrNotExist) {
		logger.Info(fmt.Sprintf("File not exists, try to create a new file. Path %s", storageFilePath))

		_, err := os.Stat("tmp")

		if err != nil && errors.Is(err, os.ErrNotExist) {
			err := os.Mkdir("tmp", 0o700)
			if err != nil {
				logger.Error(fmt.Errorf("can't create dir for storage, err %w", err))
				return
			}
		}

		file, err := os.Create(storageFilePath)
		if err != nil {
			logger.Error(fmt.Errorf("can't create file for storage, err %w", err))
			return
		}

		defer file.Close()

		logger.Info("File has been created")

		return
	}

	if err != nil {
		logger.Error(fmt.Errorf("can't check stat file for storage, err %w", err))
		return
	}

	file, err := os.Open(storageFilePath)
	if err != nil {
		logger.Error(fmt.Errorf("can't open file for storage, err %w", err))
		return
	}

	defer file.Close()

	fileData, err := io.ReadAll(file)
	if err != nil {
		logger.Error(fmt.Errorf("can't read data from file, err %w", err))
		return
	}

	if len(fileData) == 0 {
		return
	}

	listJSONURLFromFile := strings.Split(string(fileData), "\n")

	for _, data := range listJSONURLFromFile {
		// Split \n write last element like ""
		if data == "" {
			continue
		}

		var URLFromFile models.URLFromFile
		err := json.Unmarshal([]byte(data), &URLFromFile)
		if err != nil {
			logger.Error(fmt.Errorf("can't unmarshal JSON from file, err %w", err))
			continue
		}

		storage.URLRepository.AddEntity(&repository.URLEntity{ID: URLFromFile.ShortURL, URL: URLFromFile.OriginURL})
	}

	logger.Info("Init from file has been finished successfully")
}
