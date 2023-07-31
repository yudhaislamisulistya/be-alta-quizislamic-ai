package database

import (
	"errors"
	"project/config"
	"project/model"
	"strconv"
)

func GetSortQuizzes(quiz *[]model.Quiz, sortBy string, order string) (interface{}, error) {
	result := config.DB.Order(sortBy + " " + order).Find(&quiz)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return quiz, nil
}

func GetPaginationQuizzes(quiz *[]model.Quiz, page string, limit string) (interface{}, error) {
	pageInt, errPageInt := strconv.Atoi(page)
	if errPageInt != nil {
		return nil, errors.New("page must be integer")
	}

	limitInt, errLimitInt := strconv.Atoi(limit)
	if errLimitInt != nil {
		return nil, errors.New("limit must be integer")
	}

	offset := (pageInt - 1) * limitInt
	result := config.DB.Offset(offset).Limit(limitInt).Find(&quiz)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return quiz, nil
}

func GetSearchQuizzes(quiz *[]model.Quiz, keyword string) (interface{}, error) {
	result := config.DB.Where("name LIKE ?", "%"+keyword+"%").Or("description LIKE ?", "%"+keyword+"%").Or("topic LIKE ?", "%"+keyword+"%").Find(&quiz)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return quiz, nil
}
