package main

import (
	"context"
	"shorturl/internal/bootstrap"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	bootstrap.RunApp(ctx)
}
