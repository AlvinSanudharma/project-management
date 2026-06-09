package utils

import (
	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func FormatValidationError(err error) string {

	for _, e := range err.(validator.ValidationErrors) {
		switch e.Tag() {
		case "required":
			return e.Field() + " wajib diisi"
		}
	}

	return err.Error()
}
