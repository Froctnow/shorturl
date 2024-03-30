package middleware

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	httpmodels "shorturl/internal/app/httpserver/models"
	"shorturl/internal/app/log"
	"strings"
)

func DecompressMiddleware(logger log.LogClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		headerContentEncoding := c.GetHeader("Content-Encoding")

		if !strings.Contains(headerContentEncoding, "gzip") {
			c.Next()
			return
		}

		headerContentType := c.GetHeader("Content-Type")
		isCorrectContentType := checkHeaderContentType(headerContentType)

		if !isCorrectContentType {
			c.Next()
			return
		}

		gzipReader, err := gzip.NewReader(c.Request.Body)

		if err != nil {
			logger.Error(fmt.Errorf("can't get gzip reader: %w", err))
			c.AbortWithStatusJSON(http.StatusInternalServerError, httpmodels.ErrorResponse{Error: "can't get gzip reader"})
			return
		}

		defer gzipReader.Close()

		var decompressedData bytes.Buffer
		_, err = decompressedData.ReadFrom(gzipReader)
		if err != nil {
			logger.Error(fmt.Errorf("failed decompress data: %w", err))
			c.AbortWithStatusJSON(http.StatusInternalServerError, httpmodels.ErrorResponse{Error: "failed decompress data"})
			return
		}

		c.Request.Body = io.NopCloser(&decompressedData)
		c.Next()
	}
}

func checkHeaderContentType(value string) bool {
	isApplicationGzip := strings.Contains(value, "application/x-gzip")
	isTextHTML := strings.Contains(value, "text/html")

	return isApplicationGzip || isTextHTML
}
