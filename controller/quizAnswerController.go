package controller

import (
	"net/http"
	"project/lib/database"
	"project/model"

	"github.com/labstack/echo/v4"
)

func GetQuizAnswersController(c echo.Context) error {
	quizAnswers := []model.QuizAnswer{}

	result, err := database.GetQuizAnswers(&quizAnswers)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusNotFound, map[string]string{
			"code":    "404",
			"message": "Data Not Found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success get quiz",
		"data":    result,
	})
}

func GetQuizAnswerController(c echo.Context) error {
	id := c.Param("id")
	quizHistory := model.QuizAnswer{}

	result, err := database.GetQuizAnswer(&quizHistory, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to get quiz answers",
			"error":   err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    http.StatusNotFound,
			"message": "Quiz answer not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success to get quiz answers",
		"data":    quizHistory,
	})
}

func GetByQuizHistoryIDQuizAnswersController(c echo.Context) error {
	id := c.Param("id")
	quizAnswers := []model.QuizAnswer{}

	result, err := database.GetByQuizHistoryIDQuizAnswers(&quizAnswers, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to get quiz answers",
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
		"message": "Success to get quiz answers",
		"data":    quizAnswers,
	})
}

func GetByQuestionIDQuizAnswersController(c echo.Context) error {
	id := c.Param("id")
	quizAnswers := []model.QuizAnswer{}

	result, err := database.GetByQuestionIDQuizAnswers(&quizAnswers, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to get quiz answers",
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
		"message": "Success to get quiz answers",
		"data":    quizAnswers,
	})
}

func CreateQuizAnswerController(c echo.Context) error {
	quizHistory := model.QuizAnswer{}

	c.Bind(&quizHistory)

	result, err := database.CreateQuizAnswer(&quizHistory)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to create quiz answers",
			"error":   err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Failed to create quiz answers",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success to create quiz answers",
		"data":    quizHistory,
	})
}

func UpdateQuizAnswerController(c echo.Context) error {
	id := c.Param("id")
	quizHistory := model.QuizAnswer{}

	c.Bind(&quizHistory)

	_, err := database.UpdateQuizAnswer(&quizHistory, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to update quiz answers",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success to update quiz answers",
	})
}

func DeleteQuizAnswerController(c echo.Context) error {
	id := c.Param("id")

	_, err := database.DeleteQuizAnswer(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to delete quiz answers",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success to delete quiz answers",
	})
}

func GetSortQuizAnswersController(c echo.Context) error {
	sortBy := c.QueryParam("sort_by")
	order := c.QueryParam("order")
	quizAnswers := []model.QuizAnswer{}

	result, err := database.GetSortQuizAnswers(&quizAnswers, sortBy, order)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to get quiz answers",
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
		"message": "Success to get quiz answers",
		"data":    quizAnswers,
	})
}

func GetPaginationQuizAnswersController(c echo.Context) error {
	page := c.QueryParam("page")
	limit := c.QueryParam("limit")
	quizAnswers := []model.QuizAnswer{}

	result, err := database.GetPaginationQuizAnswers(&quizAnswers, page, limit)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to get quiz answers",
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
		"message": "Success to get quiz answers",
		"data":    quizAnswers,
	})
}

func GetFilterQuizAnswersController(c echo.Context) error {
	isCorrect := c.QueryParam("is_correct")
	quizAnswers := []model.QuizAnswer{}

	result, err := database.GetFilterQuizAnswers(&quizAnswers, isCorrect)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to get quiz answers",
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
		"message": "Success to get quiz answers",
		"data":    quizAnswers,
	})
}
