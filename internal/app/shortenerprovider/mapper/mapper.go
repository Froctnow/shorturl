package mapper

type mapper struct{}

type Mapper interface {
	URLIDs(urls []string) map[string]any
}

func New() Mapper {
	return &mapper{}
}
