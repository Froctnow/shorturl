package validator

import (
	"github.com/gobuffalo/validate"
	httpmodels "shorturl/internal/app/httpserver/models"
)

func (v *validator) ShortenCreateShortURL(data *httpmodels.CreateURLRequest, urlPattern string) *validate.Errors {
	checks := []validate.Validator{
		&StringLenGreaterThenValidator{
			Name:  "URL",
			Field: data.URL,
			Min:   1,
		},
		&RegexpValidator{
			Name:    "URL",
			Field:   data.URL,
			Pattern: urlPattern,
		},
	}
	errors := validate.Validate(checks...)
	return errors
}
