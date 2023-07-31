package controller

import (
	"net/http"
	"project/lib/database"
	"project/model"

	"github.com/labstack/echo/v4"
)

func GetNotificationsController(c echo.Context) error {
	notifications := []model.Notification{}

	result, err := database.GetNotifications(&notifications)

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
		"message": "success get notification",
		"data":    result,
	})
}

func GetNotificationController(c echo.Context) error {
	id := c.Param("id")
	notification := model.Notification{}

	result, err := database.GetNotification(&notification, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to get notification",
			"error":   err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    http.StatusNotFound,
			"message": "Quiz notification not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success to get notification",
		"data":    notification,
	})
}

func GetByUserIDNotificationsController(c echo.Context) error {
	id := c.Param("id")
	notifications := []model.Notification{}

	result, err := database.GetByUserIDNotifications(&notifications, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to get notification",
			"error":   err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    http.StatusNotFound,
			"message": "Quiz notifications not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success to get notification",
		"data":    notifications,
	})
}

func CreateNotificationController(c echo.Context) error {
	notification := model.Notification{}

	c.Bind(&notification)

	result, err := database.CreateNotification(&notification)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to create notification",
			"error":   err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": "Failed to create notification",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success to create notification",
		"data":    notification,
	})
}

func UpdateNotificationController(c echo.Context) error {
	id := c.Param("id")
	notification := model.Notification{}

	c.Bind(&notification)

	_, err := database.UpdateNotification(&notification, id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to update notification",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success to update notification",
	})
}

func DeleteNotificationController(c echo.Context) error {
	id := c.Param("id")

	_, err := database.DeleteNotification(id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to delete notification",
			"error":   err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success to delete notification",
	})
}

func GetSortNotificationsController(c echo.Context) error {
	sortBy := c.QueryParam("sort_by")
	order := c.QueryParam("order")
	notifications := []model.Notification{}

	result, err := database.GetSortNotifications(&notifications, sortBy, order)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to get notification",
			"error":   err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    http.StatusNotFound,
			"message": "Quiz notifications not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success to get notification",
		"data":    notifications,
	})
}

func GetPaginationNotificationsController(c echo.Context) error {
	page := c.QueryParam("page")
	limit := c.QueryParam("limit")
	notifications := []model.Notification{}

	result, err := database.GetPaginationNotifications(&notifications, page, limit)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to get notification",
			"error":   err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    http.StatusNotFound,
			"message": "Quiz notifications not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success to get notification",
		"data":    notifications,
	})
}

func GetFilterNotificationsController(c echo.Context) error {
	isCorrect := c.QueryParam("is_correct")
	notifications := []model.Notification{}

	result, err := database.GetFilterNotifications(&notifications, isCorrect)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Failed to get notification",
			"error":   err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":    http.StatusNotFound,
			"message": "Quiz notifications not found",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    http.StatusOK,
		"message": "Success to get notification",
		"data":    notifications,
	})
}
