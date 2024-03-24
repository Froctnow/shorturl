package shorten

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"net/http/httptest"
	"shorturl/internal/app/httpserver/constants"
	"shorturl/internal/app/httpserver/models"
	httpmodels "shorturl/internal/app/httpserver/models"
	"shorturl/internal/app/provider"
	"shorturl/internal/app/storage"
	"shorturl/internal/app/usecase/url"
	"shorturl/internal/app/validator"
	"testing"
)

const ServerURL = "http://localhost:8080"
const targetRoute = "/api/shorten"

func TestShortenRouter_CreateShortURL(t *testing.T) {
	ginEngine := gin.Default()

	storageMock := storage.NewStorage()
	shortenerProvider := provider.NewStorageProvider(storageMock)
	urlUseCase := url.NewUseCase(shortenerProvider, ServerURL)
	ginEngine.Use(gin.Recovery())
	validatorInstance := validator.New()

	apiGroup := ginEngine.Group("/")

	_ = NewShortenRouter(apiGroup, urlUseCase, validatorInstance)

	t.Run("incorrect content-type", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodPost, targetRoute, nil)
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
		reqData, _ := json.Marshal(httpmodels.CreateUrlRequest{URL: "some_invalid_url"})
		request := httptest.NewRequest(http.MethodPost, targetRoute, bytes.NewReader(reqData))
		request.Header.Add("Content-Type", "application/json")

		w := httptest.NewRecorder()
		ginEngine.ServeHTTP(w, request)

		result := w.Result()
		defer result.Body.Close()
		var errResponse models.ErrorResponse
		errResult, _ := io.ReadAll(result.Body)
		json.Unmarshal(errResult, &errResponse)

		validateErrorMessage := fmt.Sprintf("%s field doesn't match pattern %s", "URL", constants.RegexpURL)

		assert.Equal(t, 400, result.StatusCode)
		assert.Equal(t, validateErrorMessage, errResponse.Error)
	})

	t.Run("success convert url to short url", func(t *testing.T) {
		reqData, _ := json.Marshal(httpmodels.CreateUrlRequest{URL: "https://practicum.yandex.ru/"})
		request := httptest.NewRequest(http.MethodPost, targetRoute, bytes.NewReader(reqData))
		request.Header.Add("Content-Type", "application/json")

		w := httptest.NewRecorder()
		ginEngine.ServeHTTP(w, request)

		result := w.Result()
		defer result.Body.Close()

		assert.Equal(t, 200, result.StatusCode)
	})
}
