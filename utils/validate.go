package utils

import (
	"log"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

func FormatValidationError(err error) string {

	for _, e := range err.(validator.ValidationErrors) {
		switch e.Tag() {
		case "required":
			return e.Field() + " wajib diisi"
		case "email":
			return e.Field() + " harus berupa alamat email yang valid"
		default:
			log.Println("Unhandled validation tag:", e.Tag())
		}
	}

	return err.Error()
}
