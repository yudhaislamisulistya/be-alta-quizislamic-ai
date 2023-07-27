package controller

import (
	"fmt"
	"net/http"
	"project/lib/database"
	"project/model"

	"github.com/labstack/echo/v4"
)

func GetPackageHistoriesController(c echo.Context) error {
	packages := []model.PackageHistory{}

	result, err := database.GetPackageHistories(&packages)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": "failed get packages",
			"error":   err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    "404",
			"message": "packages not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success get packages",
		"data":    packages,
	})
}

func GetPackageHistoryController(c echo.Context) error {
	id := c.Param("id")
	packageHistories := model.PackageHistory{}

	result, err := database.GetPackageHistory(&packageHistories, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": "failed get package",
			"error":   err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    "404",
			"message": "package not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success get package",
		"data":    packageHistories,
	})
}

func CreatePackageHistoryController(c echo.Context) error {
	packageHistory := model.PackageHistory{}

	c.Bind(&packageHistory)

	result, err := database.CreatePackageHistory(&packageHistory)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": "failed create package",
			"error":   err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    "400",
			"message": "failed create package",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success create package",
		"data":    packageHistory,
	})
}

func UpdatePackageHistoryController(c echo.Context) error {
	id := c.Param("id")
	packageHistory := model.PackageHistory{}

	c.Bind(&packageHistory)

	result, err := database.UpdatePackageHistory(&packageHistory, id)
	fmt.Println(result)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": "failed update package",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success update package",
		"data":    result,
	})
}

func DeletePackageHistoryController(c echo.Context) error {
	id := c.Param("id")
	packageHistory := model.PackageHistory{}

	result, err := database.DeletePackageHistory(&packageHistory, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": "failed delete package",
			"error":   err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    "400",
			"message": "failed delete package",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success delete package",
		"data":    result,
	})
}

func GetSearchPackageHistoriesController(c echo.Context) error {
	keyword := c.QueryParam("keyword")
	packages := []model.PackageHistory{}

	result, err := database.GetSearchPackageHistories(&packages, keyword)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": "failed get search packages",
			"error":   err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    "404",
			"message": "packages not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success get search packages",
		"data":    packages,
	})
}

func GetPaginationPackageHistoriesController(c echo.Context) error {
	page := c.QueryParam("page")
	limt := c.QueryParam("limit")
	packageHistories := []model.PackageHistory{}

	result, err := database.GetPaginationPackageHistories(&packageHistories, page, limt)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": "failed get pagination package histories",
			"error":   err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    "404",
			"message": "package histories not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success get pagination package histories",
		"data":    packageHistories,
	})
}

func GetSortPackageHistoriesController(c echo.Context) error {
	sortBy := c.QueryParam("sort_by")
	order := c.QueryParam("order")
	packages := []model.PackageHistory{}

	result, err := database.GetSortPackageHistories(&packages, sortBy, order)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": "failed get sort packages",
			"error":   err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    "404",
			"message": "packages not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success get sort packages",
		"data":    packages,
	})
}

func GetByPackageIDPackageHistoriesController(c echo.Context) error {
	id := c.Param("id")
	packageHistories := []model.PackageHistory{}

	result, err := database.GetByPackageIDPackageHistories(&packageHistories, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": "failed get package histories by package id",
			"error":   err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    "404",
			"message": "package histories not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success get package histories by package id",
		"data":    packageHistories,
	})
}

func GetByUserIDPackageHistoriesController(c echo.Context) error {
	id := c.Param("id")
	packageHistories := []model.PackageHistory{}

	result, err := database.GetByUserIDPackageHistories(&packageHistories, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": "failed get package histories by user id",
			"error":   err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    "404",
			"message": "package histories not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success get package histories by user id",
		"data":    packageHistories,
	})
}

func GetFilterPackageHistoriesController(c echo.Context) error {
	status := c.QueryParam("status")
	packageHistories := []model.PackageHistory{}

	result, err := database.GetFilterPackageHistories(&packageHistories, status)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": "failed filter package histories by status",
			"error":   err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    "404",
			"message": "package histories not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success filter package histories by status",
		"data":    packageHistories,
	})
}

func GetTransactionDateRangePackageHistoriesController(c echo.Context) error {
	startDate := c.QueryParam("start_date")
	endDate := c.QueryParam("end_date")
	packageHistories := []model.PackageHistory{}

	result, err := database.GetTransactionDateRangePackageHistories(&packageHistories, startDate, endDate)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": "failed filter package histories by transaction date range",
			"error":   err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    "400",
			"message": "package histories not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success filter package histories by transaction date range",
		"data":    packageHistories,
	})
}
