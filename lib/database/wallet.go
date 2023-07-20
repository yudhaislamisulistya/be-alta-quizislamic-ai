package database

import (
	"project/config"
	"project/lib/util"
	"project/model"
)

func CreateWallet(userId uint, balance int64) (interface{}, error) {
	token := util.GetToken(64)
	wallet := model.Wallet{
		UserID:  userId,
		Balance: balance,
		Token:   token,
	}

	result := config.DB.Save(&wallet)
	err := result.Error

	if err != nil {
		return nil, err
	}

	return wallet, nil
}

func UpdateWallet(userId uint, balance int64) (interface{}, error) {
	wallet := model.Wallet{}
	result := config.DB.Where("user_id = ?", userId).First(&wallet)
	err := result.Error

	if err != nil {
		return nil, err
	}

	wallet.Balance = balance

	result = config.DB.Save(&wallet)
	err = result.Error

	if err != nil {
		return nil, err
	}

	return wallet, nil
}

func GetWallet(userId uint) (interface{}, error) {
	wallet := model.Wallet{}
	result := config.DB.Where("user_id = ?", userId).First(&wallet)
	err := result.Error

	if err != nil {
		return nil, err
	}

	return wallet, nil
}
