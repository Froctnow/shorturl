package validator

import (
	httpmodels "shorturl/internal/app/httpserver/models"

	"github.com/gobuffalo/validate"
)

type validator struct{}

type Validator interface {
	ShortenCreateShortURL(request *httpmodels.CreateURLRequest, urlPattern string) *validate.Errors
	ShortenCreateBatchShortURL(request *[]httpmodels.CreateBatchURLRequest) *validate.Errors
}

func New() Validator {
	return &validator{}
}
