package validator

import (
	"github.com/gobuffalo/validate"
	httpmodels "shorturl/internal/app/httpserver/models"
)

type validator struct{}

type Validator interface {
	ShortenCreateShortURL(request *httpmodels.CreateURLRequest, urlPattern string) *validate.Errors
}

func New() Validator {
	return &validator{}
}
