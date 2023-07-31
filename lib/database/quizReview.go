package database

import (
	"errors"
	"project/config"
	"project/model"
	"strconv"
)

func GetQuizReviews(quizReviews *[]model.QuizReview) (interface{}, error) {
	result := config.DB.Find(&quizReviews)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return quizReviews, nil
}

func GetQuizReview(QuizReview *model.QuizReview, id string) (interface{}, error) {
	idInt, errIdInt := strconv.Atoi(id)
	if errIdInt != nil {
		return nil, errors.New("id must be integer")
	}

	result := config.DB.First(QuizReview, idInt)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return QuizReview, nil
}

func GetByUserIDQuizReviews(quizReviews *[]model.QuizReview, id string) (interface{}, error) {
	idInt, errIdInt := strconv.Atoi(id)
	if errIdInt != nil {
		return nil, errors.New("id must be integer")
	}

	result := config.DB.Where("user_id = ?", idInt).Find(quizReviews)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return quizReviews, nil
}

func GetByQuizIDQuizReviews(quizReviews *[]model.QuizReview, id string) (interface{}, error) {
	idInt, errIdInt := strconv.Atoi(id)
	if errIdInt != nil {
		return nil, errors.New("id must be integer")
	}

	result := config.DB.Where("quiz_id = ?", idInt).Find(quizReviews)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return quizReviews, nil
}

func CreateQuizReview(QuizReview *model.QuizReview) (interface{}, error) {
	result := config.DB.Create(QuizReview)
	err := result.Error

	if err != nil {
		return nil, err
	}

	return QuizReview, nil
}

func UpdateQuizReview(QuizReview *model.QuizReview, id string) (interface{}, error) {
	idInt, errIdInt := strconv.Atoi(id)
	if errIdInt != nil {
		return nil, errors.New("id must be integer")
	}

	result := config.DB.Model(QuizReview).Where("id = ?", idInt).Updates(QuizReview)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return QuizReview, nil
}

func DeleteQuizReview(id string) (interface{}, error) {
	QuizReview := model.QuizReview{}
	idInt, errIdInt := strconv.Atoi(id)
	if errIdInt != nil {
		return nil, errors.New("id must be integer")
	}

	result := config.DB.Delete(&QuizReview, idInt)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return len, nil
}

func GetSortQuizReviews(quizReviews *[]model.QuizReview, sortBy string, order string) (interface{}, error) {
	result := config.DB.Order(sortBy + " " + order).Find(quizReviews)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}
	return quizReviews, nil
}

func GetPaginationQuizReviews(quizReviews *[]model.QuizReview, page string, limit string) (interface{}, error) {
	pageInt, errPageInt := strconv.Atoi(page)
	if errPageInt != nil {
		return nil, errors.New("page must be integer")
	}

	limitInt, errLimitInt := strconv.Atoi(limit)
	if errLimitInt != nil {
		return nil, errors.New("limit must be integer")
	}

	offset := (pageInt - 1) * limitInt
	result := config.DB.Offset(offset).Limit(limitInt).Find(quizReviews)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}
	return quizReviews, nil
}

func GetFilterQuizReviews(quizReviews *[]model.QuizReview, isCorrect string) (interface{}, error) {
	isCorrectBool, errIsCorrectBool := strconv.ParseBool(isCorrect)
	if errIsCorrectBool != nil {
		return nil, errors.New("is_correct must be boolean")
	}

	result := config.DB.Where("is_correct = ?", isCorrectBool).Find(quizReviews)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return quizReviews, nil
}

func GetSearchQuizReviews(quizReviews *[]model.QuizReview, keyword string) (interface{}, error) {
	result := config.DB.Where("review_text LIKE ?", "%"+keyword+"%").Find(quizReviews)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return quizReviews, nil
}
