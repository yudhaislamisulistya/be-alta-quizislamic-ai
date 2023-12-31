package controller

import (
	"net/http"
	"project/config"
	"project/lib/database"
	"project/lib/util"
	"project/model"

	"github.com/labstack/echo/v4"
)

func GetQuizzesController(c echo.Context) error {
	quiz := []model.Quiz{}
	result := config.DB.Find(&quiz)
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
		"code":    "200",
		"message": "success get quiz",
		"data":    quiz,
	})
}

func CreateQuizController(c echo.Context) error {
	quiz := model.Quiz{}
	c.Bind(&quiz)

	quiz.Token = util.GetToken(10)

	result := config.DB.Create(&quiz)
	err := result.Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success create quiz",
		"data":    quiz,
	})
}

func GetQuizController(c echo.Context) error {
	quiz := model.Quiz{}
	id := c.Param("id")

	result := config.DB.First(&quiz, id)
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
		"code":    "200",
		"message": "success get quiz",
		"data":    quiz,
	})
}

func UpdateQuizController(c echo.Context) error {
	quiz := model.Quiz{}
	id := c.Param("id")

	result := config.DB.First(&quiz, id)
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

	c.Bind(&quiz)

	result = config.DB.Save(&quiz)
	err = result.Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success update quiz",
		"data":    quiz,
	})
}

func DeleteQuizController(c echo.Context) error {
	quiz := model.Quiz{}
	id := c.Param("id")

	result := config.DB.First(&quiz, id)
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

	result = config.DB.Delete(&quiz, id)
	err = result.Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": err.Error(),
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success delete quiz",
		"data":    quiz,
	})
}

func GetByUserIDQuizController(c echo.Context) error {
	quiz := model.Quiz{}
	userId := c.Param("user_id")

	result := config.DB.Where("user_id = ?", userId).Find(&quiz)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": err.Error(),
		})
	}

	if len == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    "404",
			"message": "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success get quiz",
		"data":    quiz,
	})
}

func GetByUserIDUserQuizQuizController(c echo.Context) error {
	quiz := model.Quiz{}
	userId := c.Param("user_id")
	quizId := c.Param("quiz_id")

	result := config.DB.Where("user_id = ? AND id = ?", userId, quizId).Find(&quiz)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": err.Error(),
		})
	}

	if len == 0 {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    "404",
			"message": "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success get quiz",
		"data":    quiz,
	})
}

func GetPaginationQuizzesController(c echo.Context) error {
	page := c.QueryParam("page")
	limit := c.QueryParam("limit")
	quiz := []model.Quiz{}

	result, err := database.GetPaginationQuizzes(&quiz, page, limit)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    "404",
			"message": "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success get pagination quiz",
		"data":    quiz,
	})
}

func GetSortQuizzesController(c echo.Context) error {
	sortBy := c.QueryParam("sort_by")
	order := c.QueryParam("order")
	quiz := []model.Quiz{}

	result, err := database.GetSortQuizzes(&quiz, sortBy, order)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    "404",
			"message": "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success get sort quiz",
		"data":    quiz,
	})
}

func GetSearchQuizzesController(c echo.Context) error {
	keyword := c.QueryParam("keyword")
	quiz := []model.Quiz{}

	result, err := database.GetSearchQuizzes(&quiz, keyword)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"code":    "500",
			"message": err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    "404",
			"message": "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success get search quiz",
		"data":    quiz,
	})
}
