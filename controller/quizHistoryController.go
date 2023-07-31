package controller

import (
	"net/http"
	"project/lib/database"
	"project/lib/util"
	"project/model"

	"github.com/labstack/echo/v4"
)

func GetQuizHistoriesController(c echo.Context) error {
	quizHistories := []model.QuizHistory{}

	result, err := database.GetQuizHistories(&quizHistories)

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
		"data":    quizHistories,
	})
}

func GetQuizHistoryController(c echo.Context) error {
	id := c.Param("id")
	quizHistory := model.QuizHistory{}

	result, err := database.GetQuizHistory(&quizHistory, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to get quiz history",
			"error":   err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    http.StatusNotFound,
			"message": "Quiz history not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success to get quiz history",
		"data":    quizHistory,
	})
}

func GetByUserIDQuizHistoriesController(c echo.Context) error {
	id := c.Param("id")
	quizHitories := []model.QuizHistory{}

	result, err := database.GetByUserIDQuizHistories(&quizHitories, id)

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
		"data":    quizHitories,
	})
}

func GetByQuizIDQuizHistoriesController(c echo.Context) error {
	id := c.Param("id")
	quizHistories := []model.QuizHistory{}

	result, err := database.GetByQuizIDQuizHistories(&quizHistories, id)

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
		"data":    quizHistories,
	})
}

func CreateQuizHistoryController(c echo.Context) error {
	quizHistory := model.QuizHistory{}

	c.Bind(&quizHistory)

	// get attempt date time now and set to quiz history
	quizHistory.AttemptDate = util.GetTimeNow()

	result, err := database.CreateQuizHistory(&quizHistory)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to create quiz history",
			"error":   err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Failed to create quiz history",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success to create quiz history",
		"data":    quizHistory,
	})
}

func UpdateQuizHistoryController(c echo.Context) error {
	id := c.Param("id")
	quizHistory := model.QuizHistory{}

	c.Bind(&quizHistory)

	_, err := database.UpdateQuizHistory(&quizHistory, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to update quiz history",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success to update quiz history",
	})
}

func DeleteQuizHistoryController(c echo.Context) error {
	id := c.Param("id")

	_, err := database.DeleteQuizHistory(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to delete quiz history",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success to delete quiz history",
	})
}

func GetSortQuizHistoriesController(c echo.Context) error {
	sortBy := c.QueryParam("sort_by")
	order := c.QueryParam("order")
	quizHistories := []model.QuizHistory{}

	result, err := database.GetSortQuizHistories(&quizHistories, sortBy, order)

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
		"data":    quizHistories,
	})
}

func GetPaginationQuizHistoriesController(c echo.Context) error {
	page := c.QueryParam("page")
	limit := c.QueryParam("limit")
	quizHistories := []model.QuizHistory{}

	result, err := database.GetPaginationQuizHistories(&quizHistories, page, limit)

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
		"data":    quizHistories,
	})
}

func GetScoreQuizHistoriesController(c echo.Context) error {
	data := c.QueryParam("data")
	score := c.QueryParam("score")
	quizHistories := []model.QuizHistory{}

	result, err := database.GetScoreQuizHistories(&quizHistories, data, score)

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
		"data":    quizHistories,
	})
}

func GetAttemptDateRangeQuizHistoriesController(c echo.Context) error {
	startDate := c.QueryParam("start_date")
	endDate := c.QueryParam("end_date")
	quizHistories := []model.QuizHistory{}

	result, err := database.GetAttemptDateRangeQuizHistories(&quizHistories, startDate, endDate)

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
		"data":    quizHistories,
	})
}
