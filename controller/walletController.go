package controller

import (
	"fmt"
	"net/http"
	"project/config"
	"project/lib/database"
	"project/model"
	"strconv"

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

func SendWalletController(c echo.Context) error {
	walletSend := model.Wallet{}
	walletReceive := model.Wallet{}
	userIdSend := c.FormValue("user_id_send")
	userIdReceive := c.FormValue("user_id_receive")
	amount := c.FormValue("amount")

	resultSend := config.DB.Where("user_id = ?", userIdSend).First(&walletSend)
	errSend := resultSend.Error

	if errSend != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": "Data Pengirim Tidak Ditemukan",
		})
	}

	resultReceive := config.DB.Where("user_id = ?", userIdReceive).First(&walletReceive)
	errReceive := resultReceive.Error

	if errReceive != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": "Data Penerima Tidak Ditemukan",
		})
	}

	amountInt, err := strconv.Atoi(amount)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": err.Error(),
		})
	}

	walletSend.Balance = walletSend.Balance - int64(amountInt)

	if walletSend.Balance < 0 {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": "Saldo Tidak Mencukupi",
		})
	}

	walletReceive.Balance = walletReceive.Balance + int64(amountInt)

	resultSend = config.DB.Save(&walletSend)
	errSend = resultSend.Error

	if errSend != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": errSend.Error(),
		})
	}

	resultReceive = config.DB.Save(&walletReceive)
	errReceive = resultReceive.Error

	if errReceive != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": errReceive.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":          "200",
		"message":       "success send wallet",
		"amountSend":    walletSend.Balance,
		"walletSend":    walletSend,
		"walletReceive": walletReceive,
	})
}

func GetPaginationWalletController(c echo.Context) error {
	page := c.QueryParam("page")
	limit := c.QueryParam("limit")
	wallets := []model.Wallet{}

	result, err := database.GetPaginationWallet(&wallets, page, limit)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": err.Error(),
		})

	}

	if result == int64(0) {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    "200",
			"message": "Data Kosong",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success get pagination wallet",
		"data":    wallets,
	})
}

func GetSortWalletController(c echo.Context) error {
	sortBy := c.QueryParam("sort_by")
	order := c.QueryParam("order")
	wallets := []model.Wallet{}

	fmt.Println(sortBy)
	fmt.Println(order)

	result, err := database.GetSortWallet(&wallets, sortBy, order)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": err.Error(),
		})

	}

	if result == int64(0) {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    "200",
			"message": "Data Kosong",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success get sort wallet",
		"data":    wallets,
	})
}
