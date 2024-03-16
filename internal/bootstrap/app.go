package bootstrap

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"shorturl/internal/app/httpserver"
	"shorturl/internal/app/provider"
	"shorturl/internal/app/storage"
	"shorturl/internal/app/usecase/url"
	"syscall"
)

func RunApp(ctx context.Context) {
	ginEngine := NewGinEngine()
	httpServer, err := RunHTTPServer(ginEngine)
	if err != nil {
		panic(fmt.Errorf("http server can't start %w", err))
	}

	storageInstance := storage.NewStorage()
	storageProvider := provider.NewStorageProvider(storageInstance)
	urlUseCase := url.NewUseCase(storageProvider)

	_ = httpserver.NewShortenerServer(
		ginEngine,
		urlUseCase,
	)

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	fmt.Println("app is ready")
	select {
	case v := <-exit:
		fmt.Println((fmt.Sprintf("signal.Notify: %v", v)))
	case done := <-ctx.Done():
		fmt.Println(fmt.Errorf("ctx.Done: %v", done))
	}

	if err := httpServer.Shutdown(ctx); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Server Exited Properly")
}
