package options

type LoggerOptions struct {
	Extras    any
	Protocol  Protocol
	RequestID string
}

type Protocol string

const (
	HTTPProtocol Protocol = "http"
)

type LoggerOption func(*LoggerOptions)

func WithExtras(data any) LoggerOption {
	return func(lo *LoggerOptions) {
		lo.Extras = data
	}
}
