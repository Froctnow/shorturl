package log

import (
	"shorturl/internal/app/config"
	"shorturl/pkg/logger"
)

func New(cfg config.Values) (logger.LogClient, error) {
	log, err := logger.New(logger.Options{
		ConsoleOptions: logger.ConsoleOptions{
			Level: cfg.LogLevel,
		},
	})
	if err != nil {
		return nil, err
	}

	return log, err
}
