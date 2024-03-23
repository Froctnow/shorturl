package config

import (
	"flag"
	"fmt"
	"github.com/caarlos0/env/v6"
)

type Values struct {
	Address  string `env:"SERVER_ADDRESS" envSeparator:":"`
	Hostname string `env:"BASE_URL" envSeparator:":"`
	LogLevel string `env:"LOG_LEVEL" envSeparator:":"`
}

func NewConfig() (*Values, error) {
	var cfg Values
	address := flag.String("a", "", "address of service")
	hostname := flag.String("b", "", "hostname of service")
	logLevel := flag.String("loglevel", "", "level of logs")

	err := env.Parse(&cfg)

	if err != nil {
		panic(fmt.Errorf("can't parse env %w", err))
	}

	// разбор командной строки
	flag.Parse()

	if cfg.Address == "" {
		if *address == "" {
			*address = fmt.Sprintf(`:%d`, 8080)
		}

		cfg.Address = *address
	}
	if cfg.Hostname == "" {
		if *hostname == "" {
			*hostname = "http://localhost:8080"
		}

		cfg.Hostname = *hostname
	}

	if cfg.LogLevel == "" {
		if *logLevel == "" {
			*logLevel = "info"
		}

		cfg.LogLevel = *logLevel
	}

	return &cfg, nil
}
