package database

import (
	"errors"
	"project/config"
	"project/model"
	"strconv"
)

func GetLevels() (interface{}, error) {
	levels := []model.Level{}

	result := config.DB.Find(&levels)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return levels, nil
}

func GetLevel(level *model.Level, id string) (interface{}, error) {
	result := config.DB.Where("id = ?", id).First(&level)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return level, nil
}

func CreateLevel(level *model.Level) error {

	result := config.DB.Save(&level)
	err := result.Error

	if err != nil {
		return err
	}

	return nil
}

func UpdateLevel(level *model.Level, id string) error {

	result := config.DB.Where("id = ?", id).Updates(&level)
	err := result.Error

	if err != nil {
		return err
	}

	return nil
}

func DeleteLevel(level *model.Level, id string) error {

	result := config.DB.Unscoped().Where("id = ?", id).Delete(&level)
	err := result.Error

	if err != nil {
		return err
	}

	return nil
}

func GetSearchLevels(levels *[]model.Level, keyword string) (interface{}, error) {
	result := config.DB.Where("grade LIKE ?", "%"+keyword+"%").Find(&levels)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return levels, nil
}

func GetPaginationLevels(levels *[]model.Level, page string, limit string) (interface{}, error) {
	pageInt, errPageInt := strconv.Atoi(page)
	if errPageInt != nil {
		return nil, errors.New("page must be a number")
	}

	limitInt, errLimitInt := strconv.Atoi(limit)
	if errLimitInt != nil {
		return nil, errors.New("limit must be a number")
	}

	offset := (pageInt - 1) * limitInt
	result := config.DB.Offset(offset).Limit(limitInt).Find(&levels)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return levels, nil
}

func GetSortLevels(levsls *[]model.Level, sortBy string, order string) (interface{}, error) {
	result := config.DB.Order(sortBy + " " + order).Find(&levsls)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return levsls, nil
}
