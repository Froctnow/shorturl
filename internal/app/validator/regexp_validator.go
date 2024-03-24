package validator

import (
	"fmt"
	"github.com/gobuffalo/validate"
	"regexp"
	"strings"
)

type RegexpValidator struct {
	Name    string
	Field   string
	Message string
	Pattern string
}

func (v *RegexpValidator) IsValid(errors *validate.Errors) {
	isMatched, err := regexp.MatchString(v.Pattern, v.Field)

	if err != nil {
		errors.Add(v.Name, err.Error())
		return
	}

	if v.Message == "" {
		v.Message = fmt.Sprintf("%s field doesn't match pattern %s", v.Name, v.Pattern)
	}
	if !isMatched {
		errors.Add(strings.ToLower(v.Name), v.Message)
	}
}
