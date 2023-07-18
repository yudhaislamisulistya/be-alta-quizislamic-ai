package util

import (
	"project/config"
	"project/model"

	"github.com/google/uuid"
)

func GenerateUUID() uuid.UUID {
	result := uuid.New()
	return result
}

func GetUserControllerByUUID(uuid string) interface{} {
	user := model.User{}

	result := config.DB.Where("uuid = ?", uuid).First(&user)
	err := result.Error

	if err != nil {
		if err.Error() == "record not found" {
			return map[string]string{
				"code":    "200",
				"message": "Data Kosong",
			}
		}
		return map[string]string{
			"code":    "500",
			"message": err.Error(),
		}
	}

	return map[string]interface{}{
		"code":    "200",
		"message": "success get user",
		"data":    user,
	}
}
