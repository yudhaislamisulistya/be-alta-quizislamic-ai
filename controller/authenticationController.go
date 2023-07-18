package controller

import (
	"net/http"
	"project/config"
	"project/lib/util"
	"project/middleware"
	"project/model"

	"github.com/labstack/echo/v4"
)

func LoginUserController(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)
	password := user.Password

	result := config.DB.Where("email = ?", user.Email).First(&user)
	err := result.Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": err.Error(),
		})
	}

	err = util.VerifyPassword(user.Password, password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": err.Error(),
		})
	}

	token, exp, err := middleware.CreateToken(int(user.ID), user.UUID.String(), user.Email)

	user.Token = token
	user.TokenExpired = exp

	// update token
	config.DB.Save(&user)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login",
		"user":    user,
	})
}
