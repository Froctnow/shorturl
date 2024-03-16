package config

import (
	"flag"
	"fmt"
)

const (
	AppBuildRelease = "release"
)

type Values struct {
	Address  string
	Hostname string
}

func NewConfig() (*Values, error) {
	address := flag.String("a", "", "address of service")
	hostname := flag.String("b", "", "hostname of service")

	if *address == "" {
		*address = fmt.Sprintf(`:%d`, 8080)
	}
	if *hostname == "" {
		*hostname = "http://localhost:8080"
	}

	// разбор командной строки
	flag.Parse()
	cfg := &Values{
		Address:  *address,
		Hostname: *hostname,
	}

	fmt.Println(cfg)

	return cfg, nil
}
