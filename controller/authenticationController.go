package controller

import (
	"bytes"
	"net/http"
	"project/config"
	"project/lib/util"
	"project/middleware"
	"project/model"

	"github.com/labstack/echo/v4"
)

func LoginAuthenticationController(c echo.Context) error {
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
	user.LastLogin = util.GetTimeNow()

	config.DB.Save(&user)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success login",
		"user":    user,
	})
}

func ForgotPasswordController(c echo.Context) error {
	email := c.FormValue("email")
	user := model.User{}
	result := config.DB.Where("email = ?", email).First(&user)
	errResult := result.Error

	if errResult != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": errResult.Error(),
		})
	}

	subject := "Verification Code for Your Account - QuizIslamicAI"

	verificationCode := util.GenerateVerificationCode()

	forgotPassword := model.ForgotPassword{}
	forgotPassword.UserID = user.ID
	forgotPassword.VerificationCode = verificationCode
	forgotPassword.ExpiredAtVerification = util.GetExpiredTime(10)
	resultSaveForgotPassword := config.DB.Save(&forgotPassword)
	errSaveForgotPassword := resultSaveForgotPassword.Error

	if errSaveForgotPassword != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": errSaveForgotPassword.Error(),
		})
	}

	template, errReadTemplate := util.ReadEmailTemplateVericationCode()

	if errReadTemplate != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": errReadTemplate.Error(),
		})
	}

	data := struct {
		VerificationCode string
	}{
		VerificationCode: verificationCode,
	}

	body := new(bytes.Buffer)
	errExecuteTemplate := template.Execute(body, data)
	if errExecuteTemplate != nil {
		return errExecuteTemplate
	}

	err := util.SendMail(email, subject, body.String())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": err.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":          "success forgot password - send verification code to email",
		"email":            email,
		"verificationCode": verificationCode,
	})
}

func ChangePasswordController(c echo.Context) error {
	verificationCode := c.FormValue("verification_code")
	email := c.FormValue("email")

	user := model.User{}
	resultUser := config.DB.Where("email = ?", email).First(&user)
	errResultUser := resultUser.Error

	if errResultUser != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": errResultUser.Error(),
		})
	}

	forgotPassword := model.ForgotPassword{}
	result := config.DB.Where("verification_code = ?", verificationCode).First(&forgotPassword)
	errResult := result.Error

	if forgotPassword.UserID != user.ID {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": "verification code is not valid",
		})
	}

	if forgotPassword.IsUsedVerificationCode {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": "verification code is used",
		})
	}

	if errResult != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": errResult.Error(),
		})
	}

	today := int(util.GetExpiredTime(0))

	if today > forgotPassword.ExpiredAtVerification {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": "verification code is expired",
		})
	}

	password := c.FormValue("password")
	confirmPassword := c.FormValue("confirm_password")

	if password != confirmPassword {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": "password and confirm password is not same",
		})
	}

	hashedPassword, errHashPassword := util.HashPassword(password)

	if errHashPassword != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": errHashPassword.Error(),
		})
	}

	user.Password = hashedPassword
	config.DB.Save(&user)

	forgotPassword.IsUsedVerificationCode = true
	config.DB.Save(&forgotPassword)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success change password",
	})
}
