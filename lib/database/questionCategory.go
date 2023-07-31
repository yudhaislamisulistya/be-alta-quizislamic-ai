package database

import (
	"errors"
	"project/config"
	"project/model"
	"strconv"
)

func GetByNameQuestionCategory(questionCategories *[]model.QuestionCategory, name string) (interface{}, error) {
	result := config.DB.Where("name = ?", name).Find(&questionCategories)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return questionCategories, nil
}

func GetSearchQuestionCategories(questionCategories *[]model.QuestionCategory, search string) (interface{}, error) {
	result := config.DB.Where("name LIKE ?", "%"+search+"%").Find(&questionCategories)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return questionCategories, nil
}

func GetSortQuestionCategories(questionCategories *[]model.QuestionCategory, by string, order string) (interface{}, error) {

	result := config.DB.Order(by + " " + order).Find(&questionCategories)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return questionCategories, nil
}

func GetPaginationQuestionCategories(questionCategories *[]model.QuestionCategory, page string, limit string) (interface{}, error) {
	pageInt, errPageInt := strconv.Atoi(page)
	if errPageInt != nil || pageInt < 1 {
		return nil, errors.New("page should be a positive integer")
	}

	limitInt, errLimitInt := strconv.Atoi(limit)
	if errLimitInt != nil || limitInt < 1 {
		return nil, errors.New("limit should be a positive integer")
	}

	offset := (pageInt - 1) * limitInt
	result := config.DB.Limit(limitInt).Offset(offset).Find(&questionCategories)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return questionCategories, nil
}
