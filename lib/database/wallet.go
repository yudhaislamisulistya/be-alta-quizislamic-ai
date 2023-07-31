package database

import (
	"errors"
	"project/config"
	"project/lib/util"
	"project/model"
	"strconv"
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

func GetPaginationWallet(wallets *[]model.Wallet, page string, limit string) (interface{}, error) {

	pageInt, errPageInt := strconv.Atoi(page)
	if errPageInt != nil {
		return nil, errors.New("page must be a number")
	}

	limitInt, errLimitInt := strconv.Atoi(limit)
	if errLimitInt != nil {
		return nil, errors.New("limit must be a number")
	}

	offset := (pageInt - 1) * limitInt
	result := config.DB.Limit(limitInt).Offset(offset).Find(&wallets)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return wallets, nil
}

func GetSortWallet(wallet *[]model.Wallet, sortBy string, order string) (interface{}, error) {
	result := config.DB.Order(sortBy + " " + order).Find(&wallet)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return wallet, nil
}
