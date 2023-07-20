package controller

import (
	"net/http"
	"project/config"
	"project/model"

	"github.com/labstack/echo/v4"
)

func GetQuizCategoriesController(c echo.Context) error {
	quizCategory := []model.QuizCategory{}
	result := config.DB.Find(&quizCategory)
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
		"data":    quizCategory,
	})
}

func GetQuizCategoryController(c echo.Context) error {
	id := c.Param("id")
	quizCategory := model.QuizCategory{}

	result := config.DB.Where("id = ?", id).First(&quizCategory)
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
		"data":    quizCategory,
	})
}

func CreateQuizCategoryController(c echo.Context) error {
	quizCategory := model.QuizCategory{}
	c.Bind(&quizCategory)

	result := config.DB.Create(&quizCategory)
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
		"data":    quizCategory,
	})
}

func UpdateQuizCategoryController(c echo.Context) error {
	id := c.Param("id")
	quizCategory := model.QuizCategory{}
	tempQuizCategory := model.QuizCategory{}

	config.DB.Where("id = ?", id).First(&tempQuizCategory)

	c.Bind(&quizCategory)

	result := config.DB.Where("id = ?", id).Updates(&quizCategory)
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
		"dataBefore": tempQuizCategory,
		"dataAfter":  quizCategory,
	})
}

func DeleteQuizCategoryController(c echo.Context) error {
	id := c.Param("id")
	quizCategory := model.QuizCategory{}

	result := config.DB.Where("id = ?", id).Delete(&quizCategory)
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
