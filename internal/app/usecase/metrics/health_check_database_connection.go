package metrics

func (u *metricsUseCase) HealthCheckDatabaseConnection() error {
	return u.provider.HealthCheckConnection()
}
