package database

import (
	"errors"
	"project/config"
	"project/model"
	"strconv"
)

func GetQuizAnswers(quizAnwers *[]model.QuizAnswer) (interface{}, error) {
	result := config.DB.Find(&quizAnwers)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return quizAnwers, nil
}

func GetQuizAnswer(quizHistory *model.QuizAnswer, id string) (interface{}, error) {
	idInt, errIdInt := strconv.Atoi(id)
	if errIdInt != nil {
		return nil, errors.New("id must be integer")
	}

	result := config.DB.First(quizHistory, idInt)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return quizHistory, nil
}

func GetByQuizHistoryIDQuizAnswers(quizAnswers *[]model.QuizAnswer, id string) (interface{}, error) {
	idInt, errIdInt := strconv.Atoi(id)
	if errIdInt != nil {
		return nil, errors.New("id must be integer")
	}

	result := config.DB.Where("quiz_history_id = ?", idInt).Find(quizAnswers)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return quizAnswers, nil
}

func GetByQuestionIDQuizAnswers(quizAnswers *[]model.QuizAnswer, id string) (interface{}, error) {
	idInt, errIdInt := strconv.Atoi(id)
	if errIdInt != nil {
		return nil, errors.New("id must be integer")
	}

	result := config.DB.Where("question_id = ?", idInt).Find(quizAnswers)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return quizAnswers, nil
}

func CreateQuizAnswer(quizHistory *model.QuizAnswer) (interface{}, error) {
	result := config.DB.Create(quizHistory)
	err := result.Error

	if err != nil {
		return nil, err
	}

	return quizHistory, nil
}

func UpdateQuizAnswer(quizHistory *model.QuizAnswer, id string) (interface{}, error) {
	idInt, errIdInt := strconv.Atoi(id)
	if errIdInt != nil {
		return nil, errors.New("id must be integer")
	}

	result := config.DB.Model(quizHistory).Where("id = ?", idInt).Updates(quizHistory)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return quizHistory, nil
}

func DeleteQuizAnswer(id string) (interface{}, error) {
	quizHistory := model.QuizAnswer{}
	idInt, errIdInt := strconv.Atoi(id)
	if errIdInt != nil {
		return nil, errors.New("id must be integer")
	}

	result := config.DB.Delete(&quizHistory, idInt)
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

func GetSortQuizAnswers(quizAnswers *[]model.QuizAnswer, sortBy string, order string) (interface{}, error) {
	result := config.DB.Order(sortBy + " " + order).Find(quizAnswers)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}
	return quizAnswers, nil
}

func GetPaginationQuizAnswers(quizAnswers *[]model.QuizAnswer, page string, limit string) (interface{}, error) {
	pageInt, errPageInt := strconv.Atoi(page)
	if errPageInt != nil {
		return nil, errors.New("page must be integer")
	}

	limitInt, errLimitInt := strconv.Atoi(limit)
	if errLimitInt != nil {
		return nil, errors.New("limit must be integer")
	}

	offset := (pageInt - 1) * limitInt
	result := config.DB.Offset(offset).Limit(limitInt).Find(quizAnswers)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}
	return quizAnswers, nil
}

func GetFilterQuizAnswers(quizAnswers *[]model.QuizAnswer, isCorrect string) (interface{}, error) {
	isCorrectBool, errIsCorrectBool := strconv.ParseBool(isCorrect)
	if errIsCorrectBool != nil {
		return nil, errors.New("is_correct must be boolean")
	}

	result := config.DB.Where("is_correct = ?", isCorrectBool).Find(quizAnswers)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return quizAnswers, nil
}
