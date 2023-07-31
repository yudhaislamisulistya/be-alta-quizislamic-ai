package controller

import (
	"net/http"
	"project/lib/database"
	"project/model"

	v "project/lib/validator"

	"github.com/labstack/echo/v4"
)

func GetLevelsController(c echo.Context) error {
	levels, err := database.GetLevels()

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": err.Error(),
		})
	}

	if levels == int64(0) {
		return echo.NewHTTPError(http.StatusNotFound, map[string]interface{}{
			"code":    "404",
			"message": "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Get Levels",
		"code":    "200",
		"data":    levels,
	})
}

func GetLevelController(c echo.Context) error {
	id := c.Param("id")
	level := model.Level{}
	result, err := database.GetLevel(&level, id)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": err.Error(),
		})
	}

	if result == int64(0) {
		return echo.NewHTTPError(http.StatusNotFound, map[string]interface{}{
			"code":    "404",
			"message": "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Get Level",
		"code":    "200",
		"data":    level,
	})
}

func CreateLevelController(c echo.Context) error {

	level := model.Level{}
	c.Bind(&level)

	_, errValidator := v.LevelValidator(level)

	if errValidator != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    "400",
			"message": errValidator.Error(),
		})
	}

	err := database.CreateLevel(&level)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "Create Level",
		"code":    "201",
		"data":    level,
	})
}

func UpdateLevelController(c echo.Context) error {

	id := c.Param("id")
	level := model.Level{}
	temp_level := model.Level{}
	c.Bind(&level)

	_, errValidator := v.LevelValidator(level)

	if errValidator != nil {
		return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
			"code":    "400",
			"message": errValidator.Error(),
		})
	}

	result, errResult := database.GetLevel(&temp_level, id)

	if errResult != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": errResult.Error(),
		})
	}

	err := database.UpdateLevel(&level, id)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":          "Update Level",
		"code":             "200",
		"dataBeforeUpdate": result,
		"dataAfterUpdate":  level,
	})
}

func DeleteLevelController(c echo.Context) error {

	id := c.Param("id")
	level := model.Level{}
	result, errResult := database.GetLevel(&level, id)

	if errResult != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": errResult.Error(),
		})
	}

	err := database.DeleteLevel(&level, id)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Delete Level",
		"code":    "200",
		"data":    result,
	})
}

func GetSearchLevelsController(c echo.Context) error {
	keyword := c.QueryParam("keyword")
	levels := []model.Level{}

	result, err := database.GetSearchLevels(&levels, keyword)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": err.Error(),
		})
	}

	if result == int64(0) {
		return echo.NewHTTPError(http.StatusNotFound, map[string]interface{}{
			"code":    "404",
			"message": "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Get Search Levels",
		"code":    "200",
		"data":    levels,
	})
}

func GetPaginationLevelsController(c echo.Context) error {
	page := c.QueryParam("page")
	limit := c.QueryParam("limit")
	levels := []model.Level{}

	result, err := database.GetPaginationLevels(&levels, page, limit)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": err.Error(),
		})

	}

	if result == int64(0) {
		return echo.NewHTTPError(http.StatusNotFound, map[string]interface{}{
			"code":    "404",
			"message": "Data Not Found",
		})

	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Get Pagination Levels",
		"code":    "200",
		"data":    levels,
	})
}

func GetSortLevelsController(c echo.Context) error {
	sortBy := c.QueryParam("sort_by")
	order := c.QueryParam("order")
	level := []model.Level{}

	result, err := database.GetSortLevels(&level, sortBy, order)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": err.Error(),
		})
	}

	if result == int64(0) {
		return echo.NewHTTPError(http.StatusNotFound, map[string]interface{}{
			"code":    "404",
			"message": "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "Get Sort Levels",
		"code":    "200",
		"data":    level,
	})
}
