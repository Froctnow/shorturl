package bootstrap

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func NewGinEngine() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	return gin.New()
}

func RunHTTPServer(
	ginEngine *gin.Engine,
) (*http.Server, error) {
	server := &http.Server{
		Addr:              fmt.Sprintf(`:%d`, 8080),
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
