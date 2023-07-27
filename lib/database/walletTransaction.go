package database

import (
	"errors"
	"project/config"
	"project/model"
	"strconv"
)

func GetWalletTransactions(walletTransactions *[]model.WalletTransaction) (interface{}, error) {
	result := config.DB.Find(&walletTransactions)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return walletTransactions, nil
}

func GetWalletTransaction(walletTransaction *model.WalletTransaction, id string) (interface{}, error) {
	result := config.DB.Where("id = ?", id).First(&walletTransaction)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return walletTransaction, nil
}

func CreateWalletTransaction(walletTransaction *model.WalletTransaction) (interface{}, error) {
	result := config.DB.Create(&walletTransaction)
	err := result.Error

	if err != nil {
		return nil, err
	}

	return walletTransaction, nil
}

func UpdateWalletTransaction(walletTransaction *model.WalletTransaction, id string) (interface{}, error) {
	idInt, errIdInt := strconv.Atoi(id)

	if errIdInt != nil {
		return nil, errors.New("id must be a number")
	}

	tempWalletTransaction := model.WalletTransaction{}

	resultCheck := config.DB.Where("id = ?", idInt).First(&tempWalletTransaction)
	errCheck := resultCheck.Error
	lenCheck := resultCheck.RowsAffected

	if errCheck != nil {
		return nil, errors.New("package not found")
	}

	if lenCheck == 0 {
		return lenCheck, nil
	}

	result := config.DB.Where("id = ?", idInt).Updates(&walletTransaction)
	err := result.Error

	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"old_data": tempWalletTransaction,
		"new_data": walletTransaction,
	}, nil
}

func DeleteWalletTransaction(walletTransaction *model.WalletTransaction, id string) (interface{}, error) {

	idInt, errIdInt := strconv.Atoi(id)

	if errIdInt != nil {
		return nil, errors.New("id must be a number")
	}

	resultCheck := config.DB.Where("id = ?", idInt).First(&walletTransaction)
	errCheck := resultCheck.Error

	if errCheck != nil {
		return nil, errors.New("package not found")
	}

	result := config.DB.Where("id = ?", idInt).Delete(&walletTransaction)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return walletTransaction, nil
}

func GetPaginationWalletTransactions(walletTransactions *[]model.WalletTransaction, page string, limit string) (interface{}, error) {
	pageInt, errPageInt := strconv.Atoi(page)
	if errPageInt != nil {
		return nil, errors.New("page must be a number")
	}

	limitInt, errLimitInt := strconv.Atoi(limit)
	if errLimitInt != nil {
		return nil, errors.New("limit must be a number")
	}

	offset := (pageInt - 1) * limitInt
	result := config.DB.Offset(offset).Limit(limitInt).Find(&walletTransactions)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return walletTransactions, nil
}

func GetSortWalletTransactions(walletTransactions *[]model.WalletTransaction, sortBy string, order string) (interface{}, error) {
	result := config.DB.Order(sortBy + " " + order).Find(&walletTransactions)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return walletTransactions, nil
}

func GetScoreWalletTransactions(walletTransactions *[]model.WalletTransaction, data string, amount string) (interface{}, error) {
	amountInt, errScoreInt := strconv.Atoi(amount)
	if errScoreInt != nil {
		return nil, errors.New("amount must be integer")
	}

	result := config.DB

	if data == "greater_than" {
		result = config.DB.Where("amount > ?", amountInt).Find(walletTransactions)
	} else if data == "less_than" {
		result = config.DB.Where("amount < ?", amountInt).Find(walletTransactions)
	} else if data == "greater_than_equal" {
		result = config.DB.Where("amount >= ?", amountInt).Find(walletTransactions)
	} else if data == "less_than_equal" {
		result = config.DB.Where("amount <= ?", amountInt).Find(walletTransactions)
	} else if data == "equal" {
		result = config.DB.Where("amount = ?", amountInt).Find(walletTransactions)
	} else if data == "not_equal" {
		result = config.DB.Where("amount != ?", amountInt).Find(walletTransactions)
	}

	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}
	return walletTransactions, nil
}

func GetTransactionDateRangeWalletTransactions(walletTransactions *[]model.WalletTransaction, startDate string, endDate string) (interface{}, error) {
	result := config.DB.Where("transaction_date BETWEEN ? AND ?", startDate, endDate).Find(walletTransactions)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}
	return walletTransactions, nil
}

func GetFilterWalletTransactions(walletTransactions *[]model.WalletTransaction, transactionType string) (interface{}, error) {
	result := config.DB.Where("transaction_type = ?", transactionType).Find(walletTransactions)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}
	return walletTransactions, nil
}

func GetByWalletIDWalletTransaction(walletTransctions *[]model.WalletTransaction, id string) (interface{}, error) {

	intId, errIntId := strconv.Atoi(id)
	if errIntId != nil {
		return nil, errors.New("id must be integer")
	}

	result := config.DB.Where("wallet_id = ?", intId).Find(walletTransctions)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}
	return walletTransctions, nil
}
