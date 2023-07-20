package controller

import (
	"net/http"
	"project/config"
	"project/model"

	"github.com/labstack/echo/v4"
)

func GetWalletsController(c echo.Context) error {
	wallets := []model.Wallet{}
	result := config.DB.Find(&wallets)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": err.Error(),
		})
	}

	if len == 0 {
		return c.JSON(http.StatusNotFound, map[string]string{
			"code":    "404",
			"message": "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Get Wallets",
		"code":    "200",
		"data":    wallets,
	})
}

func GetWalletController(c echo.Context) error {
	wallet := model.Wallet{}
	id := c.Param("id")

	result := config.DB.Where("id = ?", id).First(&wallet)
	err := result.Error

	if err != nil {
		if err.Error() == "record not found" {
			return c.JSON(http.StatusOK, map[string]string{
				"code":    "200",
				"message": "Data Kosong",
			})
		}

		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success get wallet",
		"data":    wallet,
	})
}

func CreateWalletController(c echo.Context) error {
	wallet := model.Wallet{}
	c.Bind(&wallet)

	result := config.DB.Create(&wallet)
	err := result.Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success create wallet",
		"data":    wallet,
	})
}

func UpdateWalletController(c echo.Context) error {
	wallet := model.Wallet{}
	id := c.Param("id")
	c.Bind(&wallet)

	result := config.DB.Where("id = ?", id).Updates(&wallet)
	err := result.Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success update wallet",
		"data":    wallet,
	})
}

func DeleteWalletController(c echo.Context) error {
	wallet := model.Wallet{}
	id := c.Param("id")

	result := config.DB.Where("id = ?", id).Delete(&wallet)
	err := result.Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success delete wallet",
		"data":    wallet,
	})
}
