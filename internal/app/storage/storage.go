package storage

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"shorturl/internal/app/log"
	"shorturl/internal/app/provider/models"
	"strings"
)

type Instance struct {
	URLRepository IURLRepository
}

func NewStorage(fileStoragePath string, logger log.LogClient) *Instance {
	var (
		_, b, _, _ = runtime.Caller(0)
		basepath   = filepath.Dir(b)
	)
	fullFileStoragePath := ""

	if fileStoragePath != "" {
		fullFileStoragePath = basepath + fileStoragePath
	}

	storage := &Instance{URLRepository: NewURLRepository(fullFileStoragePath)}

	if fileStoragePath != "" {
		initFromFile(fullFileStoragePath, storage, logger)
	}

	return storage
}

func initFromFile(storageFilePath string, storage *Instance, logger log.LogClient) {
	_, err := os.Stat(storageFilePath)

	if err != nil && errors.Is(err, os.ErrNotExist) {
		fmt.Println(storageFilePath)
		file, err := os.Create(storageFilePath)

		if err != nil {
			logger.Error(fmt.Errorf("can't create file for storage, err %w", err))
			return
		}

		defer file.Close()

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

		storage.URLRepository.AddEntity(&URLEntity{ID: URLFromFile.ShortURL, URL: URLFromFile.OriginURL})
	}
}
