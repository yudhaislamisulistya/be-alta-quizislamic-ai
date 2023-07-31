package validator

import (
	"project/model"

	"github.com/go-playground/validator/v10"
)

func LevelValidator(level model.Level) (interface{}, error) {
	errs := validator.New().Struct(level)

	if errs != nil {
		return nil, errs
	}

	return level, nil
}
