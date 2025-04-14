package utils

import (
	"log/slog"

	"github.com/go-playground/validator/v10"
)

func Validate(st interface{}) error {
	validate := validator.New()

	err := validate.Struct(st)
	if err != nil {
		slog.Error("Failed to validate user data")
	}
	return err
}
