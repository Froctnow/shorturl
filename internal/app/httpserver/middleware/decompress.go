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
)

func DecompressMiddleware(logger log.LogClient) gin.HandlerFunc {
	return func(c *gin.Context) {
		headerContentEncoding := c.GetHeader("Content-Encoding")

		if headerContentEncoding != "gzip" {
			c.Next()
			return
		}

		headerContentType := c.GetHeader("Content-Type")

		if headerContentType != "application/json" && headerContentType != "text/html" {
			c.Next()
			return
		}

		gzipReader, err := gzip.NewReader(c.Request.Body)
		defer gzipReader.Close()

		var decompressedData bytes.Buffer
		_, err = decompressedData.ReadFrom(gzipReader)
		if err != nil {
			logger.Error(fmt.Errorf("failed decompress data: %w", err))
			c.AbortWithStatusJSON(http.StatusInternalServerError, httpmodels.ErrorResponse{Error: fmt.Sprintf("failed decompress data")})
			return
		}

		c.Request.Body = io.NopCloser(&decompressedData)
	}
}
