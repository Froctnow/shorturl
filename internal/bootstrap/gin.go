package bootstrap

import (
	"fmt"
	"net/http"
	"time"

	"shorturl/internal/app/config"

	"github.com/gin-gonic/gin"
)

func NewGinEngine() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	return gin.New()
}

func RunHTTPServer(
	ginEngine *gin.Engine,
	cfg *config.Values,
) (*http.Server, error) {
	server := &http.Server{
		Addr:              cfg.Address,
		Handler:           ginEngine,
		ReadHeaderTimeout: 10 * time.Second,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Println(err)
		}
	}()

	return server, nil
}
