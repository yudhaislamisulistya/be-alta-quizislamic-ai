package controller

import (
	"net/http"
	"project/config"
	"project/lib/database"
	"project/model"

	"github.com/labstack/echo/v4"
)

func GetQuestionCategoriesController(c echo.Context) error {
	questionCategory := []model.QuestionCategory{}
	result := config.DB.Find(&questionCategory)
	err := result.Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success",
		"data":    questionCategory,
	})
}

func GetQuestionCategoryController(c echo.Context) error {
	id := c.Param("id")
	questionCategory := model.QuestionCategory{}

	result := config.DB.Where("id = ?", id).First(&questionCategory)
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
		"message": "success",
		"data":    questionCategory,
	})
}

func CreateQuestionCategoryController(c echo.Context) error {
	questionCategory := model.QuestionCategory{}
	c.Bind(&questionCategory)

	result := config.DB.Create(&questionCategory)
	err := result.Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusCreated, map[string]interface{}{
		"code":    "201",
		"message": "success",
		"data":    questionCategory,
	})
}

func UpdateQuestionCategoryController(c echo.Context) error {
	id := c.Param("id")
	questionCategory := model.QuestionCategory{}
	tempQuestionCategory := model.QuestionCategory{}

	config.DB.Where("id = ?", id).First(&tempQuestionCategory)

	c.Bind(&questionCategory)

	result := config.DB.Where("id = ?", id).Updates(&questionCategory)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": err.Error(),
		})
	}

	if len == 0 {
		return c.JSON(http.StatusOK, map[string]string{
			"code":    "200",
			"message": "Data Kosong",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":       "200",
		"message":    "success",
		"dataBefore": tempQuestionCategory,
		"dataAfter":  questionCategory,
	})
}

func DeleteQuestionCategoryController(c echo.Context) error {
	id := c.Param("id")
	questionCategory := model.QuestionCategory{}

	result := config.DB.Where("id = ?", id).Delete(&questionCategory)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": err.Error(),
		})
	}

	if len == 0 {
		return c.JSON(http.StatusOK, map[string]string{
			"code":    "200",
			"message": "Data Kosong",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success",
	})
}

func GetByNameQuestionCategoryController(c echo.Context) error {
	name := c.Param("name")
	questionCategories := []model.QuestionCategory{}

	result, err := database.GetByNameQuestionCategory(&questionCategories, name)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusOK, map[string]string{
			"code":    "200",
			"message": "Data Kosong",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success",
		"data":    questionCategories,
	})
}

func GetSearchQuestionCategoriesController(c echo.Context) error {
	search := c.QueryParam("keyword")
	questionCategories := []model.QuestionCategory{}

	result, err := database.GetSearchQuestionCategories(&questionCategories, search)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusOK, map[string]string{
			"code":    "200",
			"message": "Data Kosong",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success",
		"data":    questionCategories,
	})
}

func GetSortQuestionCategoriesController(c echo.Context) error {
	by := c.QueryParam("sort_by")
	order := c.QueryParam("order")
	questionCatgeories := []model.QuestionCategory{}

	result, err := database.GetSortQuestionCategories(&questionCatgeories, by, order)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusOK, map[string]string{
			"code":    "200",
			"message": "Data Kosong",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success",
		"data":    questionCatgeories,
	})
}

func GetPaginationQuestionCategoriesController(c echo.Context) error {
	page := c.QueryParam("page")
	limit := c.QueryParam("limit")
	questionCategories := []model.QuestionCategory{}

	result, err := database.GetPaginationQuestionCategories(&questionCategories, page, limit)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
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
		"message": "success",
		"data":    questionCategories,
	})
}
