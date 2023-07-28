package controller

import (
	"net/http"
	"project/lib/database"
	"project/model"

	"github.com/labstack/echo/v4"
)

func GetQuizReviewsController(c echo.Context) error {
	quizReviews := []model.QuizReview{}

	result, err := database.GetQuizReviews(&quizReviews)

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

func GetQuizReviewController(c echo.Context) error {
	id := c.Param("id")
	quizReview := model.QuizReview{}

	result, err := database.GetQuizReview(&quizReview, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to get quiz reviews",
			"error":   err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    http.StatusNotFound,
			"message": "Quiz review not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success to get quiz reviews",
		"data":    quizReview,
	})
}

func GetByUserIDQuizReviewsController(c echo.Context) error {
	id := c.Param("id")
	quizReviews := []model.QuizReview{}

	result, err := database.GetByUserIDQuizReviews(&quizReviews, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to get quiz reviews",
			"error":   err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    http.StatusNotFound,
			"message": "Quiz reviews not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success to get quiz reviews",
		"data":    quizReviews,
	})
}

func GetByQuizIDQuizReviewsController(c echo.Context) error {
	id := c.Param("id")
	quizReviews := []model.QuizReview{}

	result, err := database.GetByQuizIDQuizReviews(&quizReviews, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to get quiz reviews",
			"error":   err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    http.StatusNotFound,
			"message": "Quiz reviews not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success to get quiz reviews",
		"data":    quizReviews,
	})
}

func CreateQuizReviewController(c echo.Context) error {
	quizReview := model.QuizReview{}

	c.Bind(&quizReview)

	result, err := database.CreateQuizReview(&quizReview)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to create quiz reviews",
			"error":   err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Failed to create quiz reviews",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success to create quiz reviews",
		"data":    quizReview,
	})
}

func UpdateQuizReviewController(c echo.Context) error {
	id := c.Param("id")
	quizReview := model.QuizReview{}

	c.Bind(&quizReview)

	_, err := database.UpdateQuizReview(&quizReview, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to update quiz reviews",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success to update quiz reviews",
	})
}

func DeleteQuizReviewController(c echo.Context) error {
	id := c.Param("id")

	_, err := database.DeleteQuizReview(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to delete quiz reviews",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success to delete quiz reviews",
	})
}

func GetSortQuizReviewsController(c echo.Context) error {
	sortBy := c.QueryParam("sort_by")
	order := c.QueryParam("order")
	quizReviews := []model.QuizReview{}

	result, err := database.GetSortQuizReviews(&quizReviews, sortBy, order)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to get quiz reviews",
			"error":   err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    http.StatusNotFound,
			"message": "Quiz reviews not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success to get quiz reviews",
		"data":    quizReviews,
	})
}

func GetPaginationQuizReviewsController(c echo.Context) error {
	page := c.QueryParam("page")
	limit := c.QueryParam("limit")
	quizReviews := []model.QuizReview{}

	result, err := database.GetPaginationQuizReviews(&quizReviews, page, limit)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to get quiz reviews",
			"error":   err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    http.StatusNotFound,
			"message": "Quiz reviews not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success to get quiz reviews",
		"data":    quizReviews,
	})
}

func GetFilterQuizReviewsController(c echo.Context) error {
	isCorrect := c.QueryParam("is_correct")
	quizReviews := []model.QuizReview{}

	result, err := database.GetFilterQuizReviews(&quizReviews, isCorrect)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to get quiz reviews",
			"error":   err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    http.StatusNotFound,
			"message": "Quiz reviews not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success to get quiz reviews",
		"data":    quizReviews,
	})
}
