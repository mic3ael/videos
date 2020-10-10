package validators

import (
	"gopkg.in/go-playground/validator.v9"
	"strings"
)

func ValidateCoolTitle(field validator.FieldLevel) bool {
	return strings.Contains(strings.ToLower(field.Field().String()), "cool")
}