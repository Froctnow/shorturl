package metrics

import "fmt"

func (u *metricsUseCase) HealthCheckDatabaseConnection() error {
	if u.provider == nil {
		return fmt.Errorf("storage mode is not a database, ping is not supported")
	}

	return u.provider.HealthCheckConnection()
}
