package validator

import (
	httpmodels "shorturl/internal/app/httpserver/models"

	"github.com/gobuffalo/validate"
)

func (v *validator) ShortenCreateBatchShortURL(data *[]httpmodels.CreateBatchURLRequest) *validate.Errors {
	var checks []validate.Validator

	checks = append(checks, &ArrayNotEmpty[httpmodels.CreateBatchURLRequest]{
		Name:  "Data",
		Array: *data,
	})

	errors := validate.Validate(checks...)

	return errors
}
