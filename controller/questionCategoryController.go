package controller

import (
	"net/http"
	"project/config"
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
