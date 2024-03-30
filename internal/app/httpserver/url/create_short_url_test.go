package url

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"shorturl/internal/app/httpserver/constants"
	"shorturl/internal/app/httpserver/models"
	"shorturl/internal/app/provider"
	"shorturl/internal/app/storage"
	"shorturl/internal/app/usecase/url"
	"strings"
	"testing"
)

const ServerURL = "http://localhost:8080"

func TestUrlRouter_CreateShortURL(t *testing.T) {
	ginEngine := gin.Default()

	storageMock := storage.NewStorage("", nil)
	shortenerProvider := provider.NewStorageProvider(storageMock)
	urlUseCase := url.NewUseCase(shortenerProvider, ServerURL)
	ginEngine.Use(gin.Recovery())

	apiGroup := ginEngine.Group("/")

	_ = NewURLRouter(apiGroup, urlUseCase)

	t.Run("incorrect content-type", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodPost, "/", nil)
		w := httptest.NewRecorder()
		ginEngine.ServeHTTP(w, request)

		result := w.Result()
		defer result.Body.Close()
		var errResponse models.ErrorResponse
		errResult, _ := io.ReadAll(result.Body)
		json.Unmarshal(errResult, &errResponse)

		assert.Equal(t, 400, result.StatusCode)
		assert.Equal(t, constants.MessageErrorIncorrectContentType, errResponse.Error)
	})

	t.Run("invalid url", func(t *testing.T) {
		requestBody := strings.NewReader("some_invalid_url")
		request := httptest.NewRequest(http.MethodPost, "/", requestBody)
		request.Header.Add("Content-Type", "text/plain")

		w := httptest.NewRecorder()
		ginEngine.ServeHTTP(w, request)

		result := w.Result()
		defer result.Body.Close()
		var errResponse models.ErrorResponse
		errResult, _ := io.ReadAll(result.Body)
		json.Unmarshal(errResult, &errResponse)

		assert.Equal(t, 400, result.StatusCode)
		assert.Equal(t, constants.MessageErrorIncorrectURL, errResponse.Error)
	})

	t.Run("success convert url to short url", func(t *testing.T) {
		requestBody := strings.NewReader("https://practicum.yandex.ru/")
		request := httptest.NewRequest(http.MethodPost, "/", requestBody)
		request.Header.Add("Content-Type", "text/plain")

		w := httptest.NewRecorder()
		ginEngine.ServeHTTP(w, request)

		result := w.Result()
		defer result.Body.Close()

		assert.Equal(t, 201, result.StatusCode)
	})
}
