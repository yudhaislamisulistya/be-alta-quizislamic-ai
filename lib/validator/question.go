package validator

import (
	"project/model"

	"github.com/go-playground/validator/v10"
)

func QuestionValidator(question model.Questions) (interface{}, error) {
	errs := validator.New().Struct(question)

	if errs != nil {
		return nil, errs
	}

	return question, nil
}
