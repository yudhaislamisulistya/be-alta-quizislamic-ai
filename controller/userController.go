package controller

import (
	"fmt"
	"net/http"
	"project/config"
	"project/lib/util"
	"project/middleware"
	"project/model"
	"time"

	"github.com/labstack/echo/v4"
)

func GetUsersController(c echo.Context) error {
	users := []model.User{}
	header := model.Header{}

	c.Bind(&header)
	header.Authorization = c.Request().Header.Get("Authorization")
	header.Authorization = header.Authorization[len("Bearer "):]
	claims, errClaims := middleware.ExtractClaims(header.Authorization)

	if errClaims != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": errClaims.Error(),
		})
	}

	uuid := claims["uuid"].(string)

	resultUser := util.GetUserControllerByUUID(uuid)

	if resultUser.(map[string]interface{})["data"].(model.User).Token != header.Authorization {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": "Token Tidak Valid, Silahkan login ulang untuk mendapatkan token baru",
		})
	}

	result := config.DB.Find(&users)
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
		"message": "success get user baru sekali",
		"data":    users,
	})
}

func GetUserController(c echo.Context) error {
	id := c.Param("id")
	user := model.User{}

	result := config.DB.Where("id = ?", id).First(&user)
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
		"message": "success get user",
		"data":    user,
	})
}

func CreateUserController(c echo.Context) error {
	user := model.User{}
	now := time.Now()
	c.Bind(&user)

	hashedPassword, errPassword := util.HashPassword(user.Password)

	if errPassword != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": errPassword.Error(),
		})
	}

	user.UUID = util.GenerateUUID()

	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	joinedAt := fmt.Sprintf("%04d-%02d-%02d %02d:%02d:%02d", year, month, day, hour, minute, second)

	user.JoinedAt = joinedAt

	user.Password = hashedPassword

	result := config.DB.Save(&user)
	err := result.Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success save user",
		"data":    user,
	})
}

func UpdateUserController(c echo.Context) error {
	id := c.Param("id")
	user := model.User{}
	temp_user_update := model.User{}

	config.DB.Where("id = ?", id).First(&temp_user_update)

	c.Bind(&user)

	hashedPassword, errPassword := util.HashPassword(user.Password)

	if errPassword != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": errPassword.Error(),
		})
	}

	user.Password = hashedPassword

	result := config.DB.Where("id = ?", id).Updates(&user)
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
			"message": "ID Tidak Ditemukan",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":       "200",
		"message":    "success update user",
		"dataBefore": temp_user_update,
		"dataAfter":  user,
	})
}

func DeleteUserController(c echo.Context) error {
	id := c.Param("id")
	user := model.User{}
	temp_user_delete := model.User{}

	config.DB.Where("id = ?", id).First(&temp_user_delete)

	result := config.DB.Unscoped().Where("id = ?", id).Delete(&user)

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
			"message": "ID Tidak Ditemukan",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success delete user",
		"data":    temp_user_delete,
	})
}
