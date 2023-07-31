package controller

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
	"project/config"
	"project/lib/database"
	"project/lib/util"
	"project/model"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

func GetUsersController(c echo.Context) error {
	sort := c.FormValue("sort")
	users := []model.User{}
	_, errValidateAdminToken := util.ValidateAdminToken(c)

	if errValidateAdminToken != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": errValidateAdminToken.Error(),
		})
	}

	result := config.DB.Order("id " + sort).Find(&users)
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

	tokenVerifiedEmail := util.GetToken(32)

	user.TokenVerifiedEmail = tokenVerifiedEmail

	result := config.DB.Save(&user)
	err := result.Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": err.Error(),
		})
	}

	template, errReadTemplate := util.ReadEmailTemplateUserActivation()

	if errReadTemplate != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": errReadTemplate.Error(),
		})
	}

	verificationLink := os.Getenv("URL_APP") + "/users/verification-email?token=" + tokenVerifiedEmail

	data := struct {
		VerificationLink string
	}{
		VerificationLink: verificationLink,
	}

	body := new(bytes.Buffer)
	errExecuteTemplate := template.Execute(body, data)
	if errExecuteTemplate != nil {
		return errExecuteTemplate
	}
	subject := "User Activation For Your Account - QuizIslamicAI"

	errSendMail := util.SendMail(user.Email, subject, body.String())

	if errSendMail != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": errSendMail.Error(),
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

func VerificationEmailUserController(c echo.Context) error {
	token := c.QueryParam("token")
	user := model.User{}

	result := config.DB.Where("token_verified_email = ?", token).First(&user)
	userId := user.ID
	err := result.Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": "Token Tidak Valid",
		})
	}

	if user.IsVerifiedEmail {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": "Email Sudah Terverifikasi",
		})
	}

	user.IsVerifiedEmail = true
	result = config.DB.Save(&user)
	err = result.Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": err.Error(),
		})
	}

	resultWallet, err := database.CreateWallet(uint(userId), 50)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": "Gagal Membuat Wallet",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success verification email",
		"token":   token,
		"wallet":  resultWallet,
	})
}

func CreateVerificationEmailUserController(c echo.Context) error {
	user := model.User{}
	c.Bind(&user)

	result := config.DB.Where("email = ?", user.Email).First(&user)
	err := result.Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": err.Error(),
		})
	}

	if user.IsVerifiedEmail {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": "Email Sudah Terverifikasi",
		})
	}

	tokenVerifiedEmail := util.GetToken(32)

	user.TokenVerifiedEmail = tokenVerifiedEmail

	result = config.DB.Save(&user)
	err = result.Error

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": err.Error(),
		})
	}

	template, errReadTemplate := util.ReadEmailTemplateUserActivation()

	if errReadTemplate != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": errReadTemplate.Error(),
		})
	}

	verificationLink := os.Getenv("URL_APP") + "/users/verification-email?token=" + tokenVerifiedEmail

	data := struct {
		VerificationLink string
	}{
		VerificationLink: verificationLink,
	}

	body := new(bytes.Buffer)
	errExecuteTemplate := template.Execute(body, data)
	if errExecuteTemplate != nil {
		return errExecuteTemplate
	}
	subject := "User Activation For Your Account - QuizIslamicAI"

	errSendMail := util.SendMail(user.Email, subject, body.String())

	if errSendMail != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": errSendMail.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success verification email",
		"data":    user,
	})
}

func GetFilterUsersController(c echo.Context) error {

	_, errValidateAdminToken := util.ValidateAdminToken(c)

	if errValidateAdminToken != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": errValidateAdminToken.Error(),
		})
	}

	isAdminStr := c.QueryParam("is_admin")
	accountStatus := c.QueryParam("account_status")
	isVerifiedEmailStr := c.QueryParam("is_verified_email")

	activeQueryParams := 0

	if isAdminStr != "" {
		activeQueryParams++
	}

	if accountStatus != "" {
		activeQueryParams++
	}

	if isVerifiedEmailStr != "" {
		activeQueryParams++
	}

	if activeQueryParams > 1 {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": "Hanya Bisa Memilih Satu Filter",
		})
	}

	if activeQueryParams == 0 || activeQueryParams == 3 {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": "Harus Memilih Salah Satu Filter",
		})
	}

	var result interface{}
	var err error

	if isAdminStr != "" {
		// convert isAdmin to bool type
		isAdmin, errIsAdmin := strconv.ParseBool(isAdminStr)
		if errIsAdmin != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"code":    "500",
				"message": "isAdmin harus berupa boolean",
			})
		}

		result, err = database.GetFilterIsAdminUsers(isAdmin)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"code":    "500",
				"message": "Gagal Mendapatkan Filter User",
			})
		}

		if result == int64(0) {
			return c.JSON(http.StatusOK, map[string]string{
				"code":    "200",
				"message": "Data Kosong",
			})
		}
	}

	if accountStatus != "" {
		result, err = database.GetFilterAccountStatusUsers(accountStatus)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"code":    "500",
				"message": "Gagal Mendapatkan Filter User",
			})
		}

		if result == int64(0) {
			return c.JSON(http.StatusOK, map[string]string{
				"code":    "200",
				"message": "Data Kosong",
			})
		}
	}

	if isVerifiedEmailStr != "" {
		// convert isVerifiedEmail to bool type
		isVerifiedEmail, errIsVerifiedEmail := strconv.ParseBool(isVerifiedEmailStr)
		if errIsVerifiedEmail != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"code":    "500",
				"message": "isVerifiedEmail harus berupa boolean",
			})
		}

		result, err = database.GetFilterIsEmailVerifiedUsers(isVerifiedEmail)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"code":    "500",
				"message": "Gagal Mendapatkan Filter User",
			})
		}

		if result == int64(0) {
			return c.JSON(http.StatusOK, map[string]string{
				"code":    "200",
				"message": "Data Kosong",
			})
		}
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success get filter user",
		"data":    result,
	})
}

func UpdatePasswordUserController(c echo.Context) error {

	id := c.Param("id")

	// parse id to int
	idInt, errIdInt := strconv.Atoi(id)
	if errIdInt != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": "id harus berupa integer",
		})
	}

	reqPasswordChange := model.PasswordChangeRequest{}

	c.Bind(&reqPasswordChange)

	user := model.User{}
	resultUpdatePassword, errUpdatePassword := database.UpdatePasswordUser(&user, idInt, &reqPasswordChange)

	if errUpdatePassword != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": errUpdatePassword.Error(),
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success update password user",
		"data":    resultUpdatePassword,
	})
}

func GetByGenderUsersController(c echo.Context) error {
	gender := c.Param("gender")
	users := []model.User{}

	_, errValidateAdminToken := util.ValidateAdminToken(c)

	if errValidateAdminToken != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": errValidateAdminToken.Error(),
		})
	}

	result, err := database.GetByGenderUsers(&users, gender)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": "Gagal Mendapatkan Filter User",
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusOK, map[string]string{
			"code":    "200",
			"message": "Data Kosong",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success get data by gender",
		"data":    result,
	})
}

func GetByRegistrationMethodUserController(c echo.Context) error {
	method := c.Param("method")
	fmt.Println(method)
	users := []model.User{}

	_, errValidateAdminToken := util.ValidateAdminToken(c)

	if errValidateAdminToken != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": errValidateAdminToken.Error(),
		})
	}

	result, err := database.GetByRegistrationMethodUsers(&users, method)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": "Gagal Mendapatkan Filter User",
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusOK, map[string]string{
			"code":    "200",
			"message": "Data Kosong",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success get data by registration method",
		"data":    result,
	})
}

func GetByVerifiedEmailStatusUsersController(c echo.Context) error {
	verifiedEmailStatus := c.Param("status")
	users := []model.User{}

	_, errValidateAdminToken := util.ValidateAdminToken(c)

	if errValidateAdminToken != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": errValidateAdminToken.Error(),
		})
	}

	result, err := database.GetByVerifiedEmailStatusUsers(&users, verifiedEmailStatus)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusOK, map[string]string{
			"code":    "200",
			"message": "Data Kosong",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success get data by verified email status",
		"data":    result,
	})
}

func GetByBirthYearUsersController(c echo.Context) error {
	year := c.Param("year")
	users := []model.User{}

	_, errValidateAdminToken := util.ValidateAdminToken(c)

	if errValidateAdminToken != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": errValidateAdminToken.Error(),
		})
	}

	result, err := database.GetByBirthYearUsers(&users, year)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusOK, map[string]string{
			"code":    "200",
			"message": "Data Kosong",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success get data by birth date",
		"data":    result,
	})
}

func GetEmptyProfilePhotoUsersController(c echo.Context) error {
	users := []model.User{}

	_, errValidateAdminToken := util.ValidateAdminToken(c)

	if errValidateAdminToken != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": errValidateAdminToken.Error(),
		})
	}

	result, err := database.GetEmptyProfilePhotoUsers(&users)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": "Gagal Mendapatkan Filter User",
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusOK, map[string]string{
			"code":    "200",
			"message": "Data Kosong",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success get data empty profile photo",
		"data":    result,
	})
}

func GetTokenExpiredUsersController(c echo.Context) error {
	users := []model.User{}

	_, errValidateAdminToken := util.ValidateAdminToken(c)

	if errValidateAdminToken != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": errValidateAdminToken.Error(),
		})
	}

	result, err := database.GetTokenExpiredUsers(&users)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": "Gagal Mendapatkan Filter User",
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusOK, map[string]string{
			"code":    "200",
			"message": "Data Kosong",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success get data token expired",
		"data":    result,
	})
}

func GetTokenVerifiedEmailUsersController(c echo.Context) error {
	token := c.Param("token")
	users := []model.User{}

	_, errValidateAdminToken := util.ValidateAdminToken(c)

	if errValidateAdminToken != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": errValidateAdminToken.Error(),
		})
	}

	result, err := database.GetTokenVerifiedEmailUsers(&users, token)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": "Gagal Mendapatkan Filter User",
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusOK, map[string]string{
			"code":    "200",
			"message": "Data Kosong",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success get data token verified email",
		"data":    result,
	})
}

func GetJoinedDateRangeUsersController(c echo.Context) error {
	startDate := c.QueryParam("start_date")
	endDate := c.QueryParam("end_date")
	users := []model.User{}

	_, errValidateAdminToken := util.ValidateAdminToken(c)

	if errValidateAdminToken != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": errValidateAdminToken.Error(),
		})
	}

	result, err := database.GetJoinedDateRangeUsers(&users, startDate, endDate)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusOK, map[string]string{
			"code":    "200",
			"message": "Data Kosong",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success get data joined date range",
		"data":    result,
	})
}

func GetSearchUsersController(c echo.Context) error {
	keyword := c.QueryParam("keyword")
	users := []model.User{}

	_, errValidateAdminToken := util.ValidateAdminToken(c)

	if errValidateAdminToken != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": errValidateAdminToken.Error(),
		})
	}

	result, err := database.GetSearchUsers(&users, keyword)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusOK, map[string]string{
			"code":    "200",
			"message": "Data Kosong",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success get data search",
		"data":    result,
	})
}

func GetSortUsersController(c echo.Context) error {
	sortBy := c.QueryParam("sort_by")
	order := c.QueryParam("order")
	users := []model.User{}

	_, errValidateAdminToken := util.ValidateAdminToken(c)

	if errValidateAdminToken != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": errValidateAdminToken.Error(),
		})
	}

	result, err := database.GetSortUsers(&users, sortBy, order)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    "200",
			"message": "Data Kosong",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success get data sort",
		"data":    result,
	})

}

func GetPaginationUsersController(c echo.Context) error {
	page := c.QueryParam("page")
	limit := c.QueryParam("limit")
	users := []model.User{}

	_, errValidateAdminToken := util.ValidateAdminToken(c)

	if errValidateAdminToken != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": errValidateAdminToken.Error(),
		})
	}

	result, err := database.GetPaginationUsers(&users, page, limit)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"code":    "500",
			"message": err.Error(),
		})
	}

	if result == int64(0) {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"code":    "200",
			"message": "Data Kosong",
		})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"code":    "200",
		"message": "success get data pagination",
		"data":    result,
	})
}
