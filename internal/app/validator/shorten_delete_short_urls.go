package validator

import (
	"github.com/gobuffalo/validate"
)

func (v *validator) ShortenDeleteShortURLs(data *[]string) *validate.Errors {
	var checks []validate.Validator

	for _, url := range *data {
		checks = append(checks, &StringLenGreaterThenValidator{
			Name:  "URL",
			Field: url,
			Min:   1,
		})
	}

	errors := validate.Validate(checks...)

	return errors
}
