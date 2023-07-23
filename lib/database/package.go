package database

import (
	"errors"
	"project/config"
	"project/model"
	"strconv"
)

func GetPackages(packages *[]model.Package) (interface{}, error) {
	result := config.DB.Find(&packages)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return packages, nil
}

func GetPackage(packageData *model.Package, id string) (interface{}, error) {
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

func CreatePackage(packageData *model.Package) (interface{}, error) {
	result := config.DB.Create(&packageData)
	err := result.Error

	if err != nil {
		return nil, err
	}

	return packageData, nil
}

func UpdatePackage(packageData *model.Package, id string) (interface{}, error) {
	idInt, errIdInt := strconv.Atoi(id)

	if errIdInt != nil {
		return nil, errors.New("id must be a number")
	}

	tempPackageData := model.Package{}

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

func DeletePackage(packageData *model.Package, id string) (interface{}, error) {

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
