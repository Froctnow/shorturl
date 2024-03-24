package validator

import (
	"github.com/gobuffalo/validate"
	httpmodels "shorturl/internal/app/httpserver/models"
)

type validator struct{}

type Validator interface {
	ShortenCreateShortURL(request *httpmodels.CreateUrlRequest, urlPattern string) *validate.Errors
}

func New() Validator {
	return &validator{}
}
