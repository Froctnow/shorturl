package main

import (
	"context"
	"fmt"
	"shorturl/internal/app/config"
	"shorturl/internal/app/log"
	"shorturl/internal/bootstrap"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	cfg, err := config.NewConfig()

	if err != nil {
		panic(fmt.Errorf("config read err %w", err))
	}

	logger, err := log.New(*cfg)

	if err != nil {
		panic(err)
	}

	bootstrap.RunApp(ctx, cfg, logger)
}
