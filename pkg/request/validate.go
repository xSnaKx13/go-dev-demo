package request

import (
	"github.com/go-playground/validator/v10"
)

func IsValidate[T any](body T) error {
	validate := validator.New()
	err := validate.Struct(body)
	if err != nil {
		return err
	}
	return nil
}
