package database

import (
	"project/config"
	"project/model"
)

func GetFilterIsAdminUsers(isAdmin bool) (interface{}, error) {
	users := []model.User{}
	result := config.DB.Where("is_admin = ?", isAdmin).Find(&users)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return users, nil
}
