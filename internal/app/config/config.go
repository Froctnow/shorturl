package config

import (
	"flag"
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

	// разбор командной строки
	flag.Parse()
	cfg := &Values{
		Address:  *address,
		Hostname: *hostname,
	}

	return cfg, nil
}
