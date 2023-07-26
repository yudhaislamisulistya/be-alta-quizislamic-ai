package database

import (
	"errors"
	"project/config"
	"project/lib/util"
	"project/model"
	"strconv"
)

func GetFilterIsAdminUsers(isAdmin bool) (interface{}, error) {
	users := []model.User{}

	result := config.DB.Where("is_admin = ?", isAdmin).Find(&users)

	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return users, nil

}
func GetFilterAccountStatusUsers(accountStatus string) (interface{}, error) {
	users := []model.User{}
	result := config.DB.Where("account_status = ?", accountStatus).Find(&users)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return users, nil
}

func GetFilterIsEmailVerifiedUsers(isEmailVerified bool) (interface{}, error) {
	users := []model.User{}
	result := config.DB.Where("is_verified_email = ?", isEmailVerified).Find(&users)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return users, nil
}

func UpdatePasswordUser(user *model.User, idUser int, reqPasswordChange *model.PasswordChangeRequest) (interface{}, error) {

	// check password OldPassword same with NewPassword in model PaswordChangeRequest
	resultUserByID := config.DB.Where("id = ?", idUser).First(&user)
	errResultUserByID := resultUserByID.Error

	if errResultUserByID != nil {
		return nil, errors.New("user not found")
	}

	if reqPasswordChange.OldPassword == reqPasswordChange.NewPassword {
		return nil, errors.New("old Password same with New Password")
	}

	if reqPasswordChange.NewPassword != reqPasswordChange.ConfirmNewPassword {
		return nil, errors.New("new Password and confirm new Password does not match")
	}

	// check oldPassword same with password in database
	resultPassword := config.DB.Where("id = ?", idUser).First(&user)
	errResultPassword := resultPassword.Error

	if errResultPassword != nil {
		return nil, errResultPassword
	}

	errVerifyPassword := util.VerifyPassword(user.Password, reqPasswordChange.OldPassword)

	if errVerifyPassword != nil {
		return nil, errors.New("old password does not match")
	}

	// hash new password
	hashedPassword, errHashPassword := util.HashPassword(reqPasswordChange.NewPassword)

	if errHashPassword != nil {
		return nil, errHashPassword
	}

	// update password
	ResultUpdatePassword := config.DB.Model(&user).Where("id = ?", idUser).Update("password", hashedPassword)
	errResultUpdatePassword := ResultUpdatePassword.Error

	if errResultUpdatePassword != nil {
		return nil, errResultUpdatePassword
	}

	temp_user_after := model.User{}
	resultTempUserAfter := config.DB.Where("id = ?", idUser).First(&temp_user_after)
	errResultTempUserAfter := resultTempUserAfter.Error

	if errResultTempUserAfter != nil {
		return nil, errResultTempUserAfter
	}

	// return resultPassword and temp_user_after,
	// resultPassword is password before update
	// temp_user_after is password after update
	return map[string]interface{}{
		"user_before_update": user,
		"user_after_update":  temp_user_after,
	}, nil
}

func GetByGenderUsers(users *[]model.User, gender string) (interface{}, error) {
	result := config.DB.Where("gender = ?", gender).Find(&users)
	errResult := result.Error
	len := result.RowsAffected

	if errResult != nil {
		return nil, errResult
	}

	if len == 0 {
		return len, nil
	}

	return users, nil
}

func GetByRegistrationMethodUsers(users *[]model.User, method string) (interface{}, error) {
	result := config.DB.Where("registered_via = ?", method).Find(&users)
	errResult := result.Error
	len := result.RowsAffected

	if errResult != nil {
		return nil, errResult
	}

	if len == 0 {
		return len, nil
	}

	return users, nil
}

func GetByVerifiedEmailStatusUsers(users *[]model.User, verifiedEmailStatus string) (interface{}, error) {

	VerifiedEmailStatus, errVerifiedEmailStatus := strconv.ParseBool(verifiedEmailStatus)

	if errVerifiedEmailStatus != nil {
		return nil, errors.New("verified email status must be true or false")
	}

	result := config.DB.Where("is_verified_email = ?", VerifiedEmailStatus).Find(&users)
	errResult := result.Error
	len := result.RowsAffected

	if errResult != nil {
		return nil, errResult
	}

	if len == 0 {
		return len, nil
	}

	return users, nil
}

func GetByBirthYearUsers(users *[]model.User, year string) (interface{}, error) {

	// parse string to int
	yearInt, errParseInt := strconv.Atoi(year)

	if errParseInt != nil {
		return nil, errors.New("year must be number")
	}

	result := config.DB.Where("EXTRACT(YEAR FROM birth_date) = ?", yearInt).Find(&users)
	errResult := result.Error
	len := result.RowsAffected

	if errResult != nil {
		return nil, errResult
	}

	if len == 0 {
		return len, nil
	}

	return users, nil
}

func GetEmptyProfilePhotoUsers(users *[]model.User) (interface{}, error) {
	result := config.DB.Where("profile_photo = ?", "").Find(&users)
	errResult := result.Error
	len := result.RowsAffected

	if errResult != nil {
		return nil, errResult
	}

	if len == 0 {
		return len, nil
	}

	return users, nil
}

func GetTokenExpiredUsers(users *[]model.User) (interface{}, error) {
	result := config.DB.Where("token_expired < ?", util.GetTimeNow()).Find(&users)
	errResult := result.Error
	len := result.RowsAffected

	if errResult != nil {
		return nil, errResult
	}

	if len == 0 {
		return len, nil
	}

	return users, nil
}

func GetTokenVerifiedEmailUsers(users *[]model.User, token string) (interface{}, error) {
	result := config.DB.Where("token_verified_email = ?", token).Find(&users)
	errResult := result.Error
	len := result.RowsAffected

	if errResult != nil {
		return nil, errResult
	}

	if len == 0 {
		return len, nil
	}

	return users, nil
}

func GetJoinedDateRangeUsers(users *[]model.User, startDate string, endDate string) (interface{}, error) {

	// parse string to time
	startDateParse, errStartDateParse := util.ParseDate(startDate)

	if errStartDateParse != nil {
		return nil, errors.New("start date must be date")
	}

	endDateParse, errEndDateParse := util.ParseDate(endDate)

	if errEndDateParse != nil {
		return nil, errors.New("end date must be date")
	}

	result := config.DB.Where("joined_at BETWEEN ? AND ?", startDateParse, endDateParse).Find(&users)
	errResult := result.Error
	len := result.RowsAffected

	if errResult != nil {
		return nil, errResult
	}

	if len == 0 {
		return len, nil
	}

	return users, nil

}

func GetSearchUsers(users *[]model.User, keyword string) (interface{}, error) {
	result := config.DB.Where("name LIKE ? OR email LIKE ? OR username LIKE ?", "%"+keyword+"%", "%"+keyword+"%", "%"+keyword+"%").Find(&users)
	errResult := result.Error
	len := result.RowsAffected

	if errResult != nil {
		return nil, errResult
	}

	if len == 0 {
		return len, nil
	}

	return users, nil
}

func GetSortUsers(users *[]model.User, sortBy string, order string) (interface{}, error) {
	result := config.DB.Order(sortBy + " " + order).Find(&users)
	errResult := result.Error
	len := result.RowsAffected

	if errResult != nil {
		return nil, errResult
	}

	if len == 0 {
		return len, nil
	}

	return users, nil
}

func GetPaginationUsers(users *[]model.User, page string, limit string) (interface{}, error) {

	// parse string to int
	pageInt, errParsePage := strconv.Atoi(page)
	if errParsePage != nil {
		return nil, errors.New("page must be number")
	}

	limitInt, errParseLimit := strconv.Atoi(limit)
	if errParseLimit != nil {
		return nil, errors.New("limit must be number")
	}

	offset := (pageInt - 1) * limitInt
	result := config.DB.Limit(limitInt).Offset(offset).Find(&users)
	errResult := result.Error
	len := result.RowsAffected

	if errResult != nil {
		return nil, errResult
	}

	if len == 0 {
		return len, nil
	}

	return users, nil
}
