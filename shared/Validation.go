package shared

import "github.com/go-playground/validator/v10"

func ValidateStruct(data *struct{}, validate *validator.Validate) error {
	return validate.Struct(data)
}
