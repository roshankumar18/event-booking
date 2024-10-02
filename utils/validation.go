package utils

import "github.com/go-playground/validator/v10"

func TranslateValidationErrors(ve validator.ValidationErrors) map[string]string {

	errs := make(map[string]string)
	for _, fe := range ve {
		errs[fe.Field()] = fe.Tag()
	}
	return errs
}
