package database

import (
	"project/config"
	"project/model"
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
