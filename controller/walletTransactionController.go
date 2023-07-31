package controller

import (
	"fmt"
	"net/http"
	"project/lib/database"
	"project/model"

	"github.com/labstack/echo/v4"
)

func GetWalletTransactionsController(c echo.Context) error {
	walletTransactions := []model.WalletTransaction{}

	result, err := database.GetWalletTransactions(&walletTransactions)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": "failed get walletTransactions",
			"error":   err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    "404",
			"message": "walletTransactions not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success get walletTransactions",
		"data":    walletTransactions,
	})
}

func GetWalletTransactionController(c echo.Context) error {
	id := c.Param("id")
	walletTransaction := model.WalletTransaction{}

	result, err := database.GetWalletTransaction(&walletTransaction, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": "failed get wallet transaction",
			"error":   err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    "404",
			"message": "wallet transaction not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success get wallet transaction",
		"data":    walletTransaction,
	})
}

func CreateWalletTransactionController(c echo.Context) error {
	walletTransaction := model.WalletTransaction{}

	c.Bind(&walletTransaction)

	result, err := database.CreateWalletTransaction(&walletTransaction)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": "failed create wallet transaction",
			"error":   err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    "400",
			"message": "failed create wallet transaction",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success create wallet transaction",
		"data":    walletTransaction,
	})
}

func UpdateWalletTransactionController(c echo.Context) error {
	id := c.Param("id")
	walletTransaction := model.WalletTransaction{}

	c.Bind(&walletTransaction)

	result, err := database.UpdateWalletTransaction(&walletTransaction, id)
	fmt.Println(result)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": "failed update wallet transaction",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success update wallet transaction",
		"data":    result,
	})
}

func DeleteWalletTransactionController(c echo.Context) error {
	id := c.Param("id")
	walletTransaction := model.WalletTransaction{}

	result, err := database.DeleteWalletTransaction(&walletTransaction, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": "failed delete wallet transaction",
			"error":   err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    "400",
			"message": "failed delete wallet transaction",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success delete wallet transaction",
		"data":    result,
	})
}

func GetPaginationWalletTransactionsController(c echo.Context) error {
	page := c.QueryParam("page")
	limt := c.QueryParam("limit")
	walletTransactions := []model.WalletTransaction{}

	result, err := database.GetPaginationWalletTransactions(&walletTransactions, page, limt)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": "failed get pagination walletTransactions",
			"error":   err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    "404",
			"message": "walletTransactions not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success get pagination walletTransactions",
		"data":    walletTransactions,
	})
}

func GetSortWalletTransactionsController(c echo.Context) error {
	sortBy := c.QueryParam("sort_by")
	order := c.QueryParam("order")
	walletTransactions := []model.WalletTransaction{}

	result, err := database.GetSortWalletTransactions(&walletTransactions, sortBy, order)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": "failed get sort walletTransactions",
			"error":   err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    "404",
			"message": "walletTransactions not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success get sort walletTransactions",
		"data":    walletTransactions,
	})
}

func GetAmountWalletTransactionsController(c echo.Context) error {
	data := c.QueryParam("data")
	amount := c.QueryParam("amount")
	walletTransactions := []model.WalletTransaction{}

	fmt.Println(data)
	fmt.Println(amount)

	result, err := database.GetScoreWalletTransactions(&walletTransactions, data, amount)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to get wallet transactions",
			"error":   err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    http.StatusNotFound,
			"message": "Wallet transactions not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success to get wallet transactions",
		"data":    walletTransactions,
	})
}

func GetTransactionDateRangeWalletTransactionsController(c echo.Context) error {
	startDate := c.QueryParam("start_date")
	endDate := c.QueryParam("end_date")
	walletTransactions := []model.WalletTransaction{}

	result, err := database.GetTransactionDateRangeWalletTransactions(&walletTransactions, startDate, endDate)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to get quiz histories",
			"error":   err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    http.StatusNotFound,
			"message": "Quiz histories not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success to get quiz histories",
		"data":    walletTransactions,
	})
}

func GetFilterWalletTransactionsController(c echo.Context) error {
	transactionType := c.QueryParam("transaction_type")
	walletTransactions := []model.WalletTransaction{}

	result, err := database.GetFilterWalletTransactions(&walletTransactions, transactionType)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to get quiz histories",
			"error":   err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    http.StatusNotFound,
			"message": "Quiz histories not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success to get quiz histories",
		"data":    walletTransactions,
	})
}

func GetByWalletIDWalletTransactionController(c echo.Context) error {
	id := c.Param("id")
	walletTransactions := []model.WalletTransaction{}

	result, err := database.GetByWalletIDWalletTransaction(&walletTransactions, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to get wallet transactions",
			"error":   err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    http.StatusNotFound,
			"message": "Wallet transactions not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success to get wallet transactions",
		"data":    walletTransactions,
	})
}
