package storage

type Instance struct {
	URLRepository IURLRepository
}

func NewStorage() *Instance {
	storage := &Instance{URLRepository: NewURLRepository()}

	return storage
}
