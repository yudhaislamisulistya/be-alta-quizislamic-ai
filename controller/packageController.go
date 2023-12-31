package controller

import (
	"fmt"
	"net/http"
	"project/lib/database"
	"project/model"

	"github.com/labstack/echo/v4"
)

func GetPackagesController(c echo.Context) error {
	packages := []model.Package{}

	result, err := database.GetPackages(&packages)

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

func GetPackageController(c echo.Context) error {
	id := c.Param("id")
	packageData := model.Package{}

	result, err := database.GetPackage(&packageData, id)

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
		"data":    packageData,
	})
}

func CreatePackageController(c echo.Context) error {
	packageData := model.Package{}

	c.Bind(&packageData)

	result, err := database.CreatePackage(&packageData)

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
		"data":    packageData,
	})
}

func UpdatePackageController(c echo.Context) error {
	id := c.Param("id")
	packageData := model.Package{}

	c.Bind(&packageData)

	result, err := database.UpdatePackage(&packageData, id)
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

func DeletePackageController(c echo.Context) error {
	id := c.Param("id")
	packageData := model.Package{}

	result, err := database.DeletePackage(&packageData, id)

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

func GetSearchPackagesController(c echo.Context) error {
	keyword := c.QueryParam("keyword")
	packages := []model.Package{}

	result, err := database.GetSearchPackages(&packages, keyword)

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

func GetPaginationPackagesController(c echo.Context) error {
	page := c.QueryParam("page")
	limt := c.QueryParam("limit")
	packages := []model.Package{}

	result, err := database.GetPaginationPackages(&packages, page, limt)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": "failed get pagination packages",
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
		"message": "success get pagination packages",
		"data":    packages,
	})
}

func GetSortPackagesController(c echo.Context) error {
	sortBy := c.QueryParam("sort_by")
	order := c.QueryParam("order")
	packages := []model.Package{}

	result, err := database.GetSortPackages(&packages, sortBy, order)

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
