package storage

type InstanceMock struct {
	URLRepositoryMock IURLRepository
}

func NewStorageMock() *InstanceMock {
	storage := &InstanceMock{URLRepositoryMock: NewURLRepositoryMock()}

	return storage
}
