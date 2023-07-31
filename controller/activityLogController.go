package controller

import (
	"net/http"
	"project/lib/database"
	"project/model"

	"github.com/labstack/echo/v4"
)

func GetActivityLogsController(c echo.Context) error {
	activityLogs := []model.ActivityLog{}

	result, err := database.GetActivityLogs(&activityLogs)

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

func GetActivityLogController(c echo.Context) error {
	id := c.Param("id")
	quizReview := model.ActivityLog{}

	result, err := database.GetActivityLog(&quizReview, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to get activity log",
			"error":   err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    http.StatusNotFound,
			"message": "Quiz activity log not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success to get activity log",
		"data":    quizReview,
	})
}

func GetByUserIDActivityLogsController(c echo.Context) error {
	id := c.Param("id")
	activityLogs := []model.ActivityLog{}

	result, err := database.GetByUserIDActivityLogs(&activityLogs, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to get activity log",
			"error":   err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    http.StatusNotFound,
			"message": "Quiz activity logs not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success to get activity log",
		"data":    activityLogs,
	})
}

func CreateActivityLogController(c echo.Context) error {
	quizReview := model.ActivityLog{}

	c.Bind(&quizReview)

	result, err := database.CreateActivityLog(&quizReview)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to create activity log",
			"error":   err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Failed to create activity log",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success to create activity log",
		"data":    quizReview,
	})
}

func UpdateActivityLogController(c echo.Context) error {
	id := c.Param("id")
	quizReview := model.ActivityLog{}

	c.Bind(&quizReview)

	_, err := database.UpdateActivityLog(&quizReview, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to update activity log",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success to update activity log",
	})
}

func DeleteActivityLogController(c echo.Context) error {
	id := c.Param("id")

	_, err := database.DeleteActivityLog(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to delete activity log",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success to delete activity log",
	})
}

func GetSortActivityLogsController(c echo.Context) error {
	sortBy := c.QueryParam("sort_by")
	order := c.QueryParam("order")
	activityLogs := []model.ActivityLog{}

	result, err := database.GetSortActivityLogs(&activityLogs, sortBy, order)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to get activity log",
			"error":   err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    http.StatusNotFound,
			"message": "Quiz activity logs not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success to get activity log",
		"data":    activityLogs,
	})
}

func GetPaginationActivityLogsController(c echo.Context) error {
	page := c.QueryParam("page")
	limit := c.QueryParam("limit")
	activityLogs := []model.ActivityLog{}

	result, err := database.GetPaginationActivityLogs(&activityLogs, page, limit)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to get activity log",
			"error":   err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    http.StatusNotFound,
			"message": "Quiz activity logs not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success to get activity log",
		"data":    activityLogs,
	})
}

func GetFilterActivityLogsController(c echo.Context) error {
	isCorrect := c.QueryParam("is_correct")
	activityLogs := []model.ActivityLog{}

	result, err := database.GetFilterActivityLogs(&activityLogs, isCorrect)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to get activity log",
			"error":   err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    http.StatusNotFound,
			"message": "Quiz activity logs not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success to get activity log",
		"data":    activityLogs,
	})
}
