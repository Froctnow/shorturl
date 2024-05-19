package middleware

import (
	"bytes"
	"fmt"
	"io"
	"shorturl/pkg/logger"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func LoggingMiddleware(logger logger.LogClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		// Чтение и сохранение тела запроса
		bodyBytes, err := io.ReadAll(c.Request.Body)
		if err != nil {
			logger.Error(err)
		}

		// Восстановление тела запроса
		c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		methodName := c.Request.Method + " " + c.Request.URL.Path
		logger.Info("HTTP request on method: " + methodName + ", Body: " + string(bodyBytes))

		// Обработка запроса
		c.Next()

		// Логирование ответа
		duration := time.Since(start)
		statusCode := c.Writer.Status()

		requestInfoText := fmt.Sprintf(
			"HTTP method: %s, time spent: %v, status code: %d, size of response %d",
			methodName,
			duration,
			statusCode,
			c.Writer.Size(),
		)

		if c.Errors.Last() != nil {
			logger.Error(c.Errors.Last())
		} else {
			logger.Info(requestInfoText)
		}

		if err != nil {
			logger.Error(errors.Wrap(err, requestInfoText))
		}
	}
}
