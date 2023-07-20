package util

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"log"
	"project/config"
	"project/model"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func GenerateUUID() uuid.UUID {
	result := uuid.New()
	return result
}

func GetUserControllerByUUID(uuid string) interface{} {

	if uuid == "" {
		return map[string]string{
			"code":    "500",
			"message": "UUID Tidak Boleh Kosong",
		}
	}

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

func GenerateToken(email string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(email), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hash to store:", string(hash))

	hasher := md5.New()
	hasher.Write(hash)
	return hex.EncodeToString(hasher.Sum(nil))
}
