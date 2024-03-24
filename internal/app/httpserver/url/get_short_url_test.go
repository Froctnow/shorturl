package url

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"shorturl/internal/app/httpserver/constants"
	"shorturl/internal/app/httpserver/models"
	"shorturl/internal/app/provider"
	"shorturl/internal/app/storage"
	"shorturl/internal/app/usecase/url"
	"testing"
)

func TestUrlRouter_GetShortURL(t *testing.T) {
	ginEngine := gin.Default()

	storageMock := storage.NewStorage()
	shortenerProvider := provider.NewStorageProvider(storageMock)
	urlUseCase := url.NewUseCase(shortenerProvider, ServerURL)
	ginEngine.Use(gin.Recovery())

	apiGroup := ginEngine.Group("/")

	_ = NewURLRouter(apiGroup, urlUseCase)

	t.Run("test :id like uuid", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/not_uuid", nil)
		w := httptest.NewRecorder()
		ginEngine.ServeHTTP(w, request)

		result := w.Result()
		defer result.Body.Close()
		var errResponse models.ErrorResponse
		errResult, _ := io.ReadAll(result.Body)
		json.Unmarshal(errResult, &errResponse)

		assert.Equal(t, 400, result.StatusCode)
		assert.Equal(t, constants.MessageErrorIncorrectAlias, errResponse.Error)
	})

	t.Run("short url not found", func(t *testing.T) {
		target := "/" + uuid.New().String()
		request := httptest.NewRequest(http.MethodGet, target, nil)
		request.Header.Add("Content-Type", "text/plain")
		w := httptest.NewRecorder()
		ginEngine.ServeHTTP(w, request)

		result := w.Result()
		defer result.Body.Close()
		var errResponse models.ErrorResponse
		errResult, _ := io.ReadAll(result.Body)
		json.Unmarshal(errResult, &errResponse)

		assert.Equal(t, 404, result.StatusCode)
		assert.Equal(t, constants.MessageErrorShortURLNotFound, errResponse.Error)
	})

	t.Run("success get short url", func(t *testing.T) {
		redirectURL := "https://practicum.yandex.ru/"
		urlEntityDto := &storage.URLEntityDto{URL: redirectURL}
		urlEntity := storageMock.URLRepository.CreateEntity(urlEntityDto)
		target := "/" + urlEntity.ID

		request := httptest.NewRequest(http.MethodGet, target, nil)
		request.Header.Add("Content-Type", "text/plain")

		w := httptest.NewRecorder()
		ginEngine.ServeHTTP(w, request)

		result := w.Result()
		defer result.Body.Close()
		resultURL := result.Header.Get("Location")

		assert.Equal(t, 307, result.StatusCode)
		assert.Equal(t, redirectURL, resultURL)
	})
}
