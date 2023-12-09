package handlers

import (
	"fmt"
	"github.com/google/uuid"
	"io"
	"strings"

	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"shorturl/internal/services"
	"shorturl/internal/storage"
	"testing"
)

func TestHandleIndexPost(t *testing.T) {
	storageMock := storage.NewStorageMock()
	urlService := services.NewURLService(storageMock.URLRepositoryMock)
	urlHandler := NewURLHandler(urlService)

	t.Run("incorrect content-type", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodPost, "/", nil)
		w := httptest.NewRecorder()
		h := http.HandlerFunc(urlHandler.handleIndexPost)
		h(w, request)

		result := w.Result()
		errorMessage, _ := io.ReadAll(result.Body)

		assert.Equal(t, 400, result.StatusCode)
		assert.Equal(t, MessageErrorIncorrectContentType, strings.TrimSuffix(string(errorMessage), "\n"))
	})

	t.Run("invalid url", func(t *testing.T) {
		requestBody := strings.NewReader("some_invalid_url")
		request := httptest.NewRequest(http.MethodPost, "/", requestBody)
		request.Header.Add("Content-Type", "text/plain")

		w := httptest.NewRecorder()
		h := http.HandlerFunc(urlHandler.handleIndexPost)
		h(w, request)

		result := w.Result()
		errorMessage, _ := io.ReadAll(result.Body)

		assert.Equal(t, 400, result.StatusCode)
		assert.Equal(t, MessageErrorIncorrectURL, strings.TrimSuffix(string(errorMessage), "\n"))
	})

	t.Run("success convert url to short url", func(t *testing.T) {
		requestBody := strings.NewReader("https://practicum.yandex.ru/")
		request := httptest.NewRequest(http.MethodPost, "/", requestBody)
		request.Header.Add("Content-Type", "text/plain")

		w := httptest.NewRecorder()
		h := http.HandlerFunc(urlHandler.handleIndexPost)
		h(w, request)

		result := w.Result()
		url, _ := io.ReadAll(result.Body)

		fmt.Println(string(url))

		assert.Equal(t, 201, result.StatusCode)
	})
}

func TestHandleIndexGet(t *testing.T) {
	storageMock := storage.NewStorageMock()
	urlService := services.NewURLService(storageMock.URLRepositoryMock)
	urlHandler := NewURLHandler(urlService)

	t.Run("test :id like uuid", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/not_uuid", nil)
		w := httptest.NewRecorder()
		h := http.HandlerFunc(urlHandler.handleIndexGet)
		h(w, request)

		result := w.Result()
		errorMessage, _ := io.ReadAll(result.Body)

		assert.Equal(t, 400, result.StatusCode)
		assert.Equal(t, MessageErrorIncorrectID, strings.TrimSuffix(string(errorMessage), "\n"))
	})

	t.Run("incorrect content-type", func(t *testing.T) {
		target := "/" + uuid.New().String()
		request := httptest.NewRequest(http.MethodGet, target, nil)
		w := httptest.NewRecorder()
		h := http.HandlerFunc(urlHandler.handleIndexGet)
		h(w, request)

		result := w.Result()
		errorMessage, _ := io.ReadAll(result.Body)

		assert.Equal(t, 400, result.StatusCode)
		assert.Equal(t, MessageErrorIncorrectContentType, strings.TrimSuffix(string(errorMessage), "\n"))
	})

	t.Run("short url not found", func(t *testing.T) {
		target := "/" + uuid.New().String()
		request := httptest.NewRequest(http.MethodGet, target, nil)
		request.Header.Add("Content-Type", "text/plain")
		w := httptest.NewRecorder()
		h := http.HandlerFunc(urlHandler.handleIndexGet)
		h(w, request)

		result := w.Result()
		errorMessage, _ := io.ReadAll(result.Body)

		assert.Equal(t, 404, result.StatusCode)
		assert.Equal(t, MessageErrorShortURLNotFound, strings.TrimSuffix(string(errorMessage), "\n"))
	})

	t.Run("success get short url", func(t *testing.T) {
		redirectURL := "https://practicum.yandex.ru/"
		urlEntityDto := &storage.URLEntityDto{URL: redirectURL}
		urlEntity := storageMock.URLRepositoryMock.CreateEntity(urlEntityDto)
		target := "/" + urlEntity.ID

		request := httptest.NewRequest(http.MethodGet, target, nil)
		request.Header.Add("Content-Type", "text/plain")

		w := httptest.NewRecorder()
		h := http.HandlerFunc(urlHandler.handleIndexGet)
		h(w, request)

		result := w.Result()
		resultURL := result.Header.Get("Location")

		assert.Equal(t, 307, result.StatusCode)
		assert.Equal(t, redirectURL, resultURL)
	})
}
