package bootstrap

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"shorturl/internal/app/client/pg"
	"shorturl/internal/app/config"
	"shorturl/internal/app/httpserver"
	"shorturl/internal/app/log"
	"shorturl/internal/app/provider"
	"shorturl/internal/app/storage"
	"shorturl/internal/app/usecase/metrics"
	"shorturl/internal/app/usecase/url"
	"shorturl/internal/app/validator"
	"syscall"
)

func RunApp(ctx context.Context, cfg *config.Values, logger log.LogClient) {
	ginEngine := NewGinEngine()
	httpServer, err := RunHTTPServer(ginEngine, cfg)
	if err != nil {
		panic(fmt.Errorf("http server can't start %w", err))
	}

	shortenerDBConn, err := pg.New(cfg, logger)
	if err != nil {
		logger.Fatal(err)
	}

	shortenerProvider := provider.NewShortenerProvider(shortenerDBConn)

	storageInstance := storage.NewStorage(cfg.FileStoragePath, logger)
	storageProvider := provider.NewStorageProvider(storageInstance)
	urlUseCase := url.NewUseCase(storageProvider, cfg.Hostname)
	metricsUseCase := metrics.NewUseCase(shortenerProvider)
	validatorInstance := validator.New()

	_ = httpserver.NewShortenerServer(
		ginEngine,
		urlUseCase,
		logger,
		validatorInstance,
		metricsUseCase,
	)

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	fmt.Println("app is ready")
	select {
	case v := <-exit:
		fmt.Printf("signal.Notify: %v\n\n", v)
	case done := <-ctx.Done():
		fmt.Println(fmt.Errorf("ctx.Done: %v", done))
	}

	if err := httpServer.Shutdown(ctx); err != nil {
		fmt.Println(err)
	}

	fmt.Println("Server Exited Properly")
}
