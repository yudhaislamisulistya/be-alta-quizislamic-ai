package database

import (
	"errors"
	"project/config"
	"project/model"
	"strconv"
)

func GetPackageHistories(packageHistories *[]model.PackageHistory) (interface{}, error) {
	result := config.DB.Find(&packageHistories)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return packageHistories, nil
}

func GetPackageHistory(packageData *model.PackageHistory, id string) (interface{}, error) {
	result := config.DB.Where("id = ?", id).First(&packageData)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return packageData, nil
}

func CreatePackageHistory(packageData *model.PackageHistory) (interface{}, error) {
	result := config.DB.Create(&packageData)
	err := result.Error

	if err != nil {
		return nil, err
	}

	return packageData, nil
}

func UpdatePackageHistory(packageData *model.PackageHistory, id string) (interface{}, error) {
	idInt, errIdInt := strconv.Atoi(id)

	if errIdInt != nil {
		return nil, errors.New("id must be a number")
	}

	tempPackageData := model.PackageHistory{}

	resultCheck := config.DB.Where("id = ?", idInt).First(&tempPackageData)
	errCheck := resultCheck.Error
	lenCheck := resultCheck.RowsAffected

	if errCheck != nil {
		return nil, errors.New("package not found")
	}

	if lenCheck == 0 {
		return lenCheck, nil
	}

	result := config.DB.Where("id = ?", idInt).Updates(&packageData)
	err := result.Error

	if err != nil {
		return nil, err
	}

	return map[string]interface{}{
		"old_data": tempPackageData,
		"new_data": packageData,
	}, nil
}

func DeletePackageHistory(packageData *model.PackageHistory, id string) (interface{}, error) {

	idInt, errIdInt := strconv.Atoi(id)

	if errIdInt != nil {
		return nil, errors.New("id must be a number")
	}

	resultCheck := config.DB.Where("id = ?", idInt).First(&packageData)
	errCheck := resultCheck.Error

	if errCheck != nil {
		return nil, errors.New("package not found")
	}

	result := config.DB.Where("id = ?", idInt).Delete(&packageData)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return packageData, nil
}

func GetSearchPackageHistories(packageHistories *[]model.PackageHistory, keyword string) (interface{}, error) {
	result := config.DB.Where("status LIKE ?", "%"+keyword+"%").Find(&packageHistories)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return packageHistories, nil
}

func GetPaginationPackageHistories(packageHistories *[]model.PackageHistory, page string, limit string) (interface{}, error) {
	pageInt, errPageInt := strconv.Atoi(page)
	if errPageInt != nil {
		return nil, errors.New("page must be a number")
	}

	limitInt, errLimitInt := strconv.Atoi(limit)
	if errLimitInt != nil {
		return nil, errors.New("limit must be a number")
	}

	offset := (pageInt - 1) * limitInt
	result := config.DB.Offset(offset).Limit(limitInt).Find(&packageHistories)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return packageHistories, nil
}

func GetSortPackageHistories(packageHistories *[]model.PackageHistory, sortBy string, order string) (interface{}, error) {
	result := config.DB.Order(sortBy + " " + order).Find(&packageHistories)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return packageHistories, nil
}

func GetByPackageIDPackageHistories(packageHistories *[]model.PackageHistory, packageID string) (interface{}, error) {

	intPackageID, errIntPackageID := strconv.Atoi(packageID)
	if errIntPackageID != nil {
		return nil, errors.New("package id must be a number lala")
	}

	result := config.DB.Where("package_id = ?", intPackageID).Find(&packageHistories)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return packageHistories, nil
}

func GetByUserIDPackageHistories(packageHistories *[]model.PackageHistory, userID string) (interface{}, error) {

	intUserID, errIntUserID := strconv.Atoi(userID)
	if errIntUserID != nil {
		return nil, errors.New("user id must be a number")
	}

	result := config.DB.Where("user_id = ?", intUserID).Find(&packageHistories)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return packageHistories, nil
}

func GetFilterPackageHistories(packageHistories *[]model.PackageHistory, status string) (interface{}, error) {
	result := config.DB.Where("status = ?", status).Find(&packageHistories)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return packageHistories, nil
}

func GetTransactionDateRangePackageHistories(packageHistories *[]model.PackageHistory, startDate string, endDate string) (interface{}, error) {
	result := config.DB.Where("transaction_date BETWEEN ? AND ?", startDate, endDate).Find(&packageHistories)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}
	return packageHistories, nil
}
