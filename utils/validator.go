package utils

import (
	"fmt"
	"regexp"

	"github.com/go-playground/validator/v10"
)

var Validate = validator.New()

// InitValidator digunakan untuk register custom validation
func InitValidator() {

	// custom validator untuk password kuat
	Validate.RegisterValidation("password", PasswordValidator)
}

// PasswordValidator validasi password kuat
func PasswordValidator(fl validator.FieldLevel) bool {

	password := fl.Field().String()

	// minimal:
	// 1 uppercase
	// 1 lowercase
	// 1 number
	// 1 special character

	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)

	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)

	hasNumber := regexp.MustCompile(`[0-9]`).MatchString(password)

	hasSpecial := regexp.MustCompile(
		`[!@#~$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]`,
	).MatchString(password)

	return hasUpper &&
		hasLower &&
		hasNumber &&
		hasSpecial
}

// FormatValidationError mengubah validator error menjadi map
func FormatValidationError(err error) map[string]string {

	errors := make(map[string]string)

	validationErrors, ok := err.(validator.ValidationErrors)
	if !ok {
		errors["validation"] = "validation error"
		return errors
	}

	for _, err := range validationErrors {

		field := err.Field()

		switch err.Tag() {

		case "required":
			errors[field] = fmt.Sprintf(
				"%s is required",
				field,
			)

		case "email":
			errors[field] = fmt.Sprintf(
				"%s must be valid email",
				field,
			)

		case "min":
			errors[field] = fmt.Sprintf(
				"%s minimum length is %s",
				field,
				err.Param(),
			)

		case "max":
			errors[field] = fmt.Sprintf(
				"%s maximum length is %s",
				field,
				err.Param(),
			)

		case "password":
			errors[field] =
				"password must contain uppercase, lowercase, number, and special character"

		default:
			errors[field] = fmt.Sprintf(
				"%s is invalid",
				field,
			)
		}
	}

	return errors
}