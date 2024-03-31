package pg

import (
	"embed"
	"fmt"
	"shorturl/internal/app/config"
	"shorturl/internal/app/log"
	"shorturl/pkg/pgclient"
)

var (
	//go:embed queries/*
	queryFiles embed.FS

	pathsToDbQueries = []string{"queries/"}
)

//go:generate mockery --srcpkg=vcs.bingo-boom.ru/bb_online/go-modules/pgclient --case=underscore --name=Transaction

func New(cfg *config.Values, log log.LogClient) (pgclient.PGClient, error) {
	if cfg == nil {
		return nil, fmt.Errorf("invalid pg config")
	}
	if cfg.DatabaseDSN == "" {
		return nil, nil
	}
	connString := cfg.DatabaseDSN
	configValues := pgclient.PostgreSQL{
		ConnString:     connString,
		PathsToQueries: pathsToDbQueries,
		LogLevel:       pgclient.LogLevelNone,
	}
	return pgclient.New(configValues, log, queryFiles)
}
