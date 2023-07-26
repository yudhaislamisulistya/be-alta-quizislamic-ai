package config

import (
	"fmt"
	"log"
	"os"
	"project/model"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitDB() {

	config := model.Config{
		DB_Username: os.Getenv("DB_USER"),
		DB_Password: os.Getenv("DB_PASS"),
		DB_Port:     os.Getenv("DB_PORT"),
		DB_Host:     os.Getenv("DB_HOST"),
		DB_Name:     os.Getenv("DB_NAME"),
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DB_Username,
		config.DB_Password,
		config.DB_Host,
		config.DB_Port,
		config.DB_Name,
	)

	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

func InitialMigration() {
	DB.AutoMigrate(&model.User{})
	DB.AutoMigrate(&model.ForgotPassword{})
	DB.AutoMigrate(&model.Questions{})
	DB.AutoMigrate(&model.QuestionCategory{})
	DB.AutoMigrate(&model.Wallet{})
	DB.AutoMigrate(&model.Quiz{})
	DB.AutoMigrate(&model.Level{})
	DB.AutoMigrate(&model.Package{})
	DB.AutoMigrate(&model.QuizHistory{})
}

func LoadDotEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
