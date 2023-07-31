package database

import (
	"errors"
	"project/config"
	"project/model"
	"strconv"
)

func GetQuizHistories(quizHistories *[]model.QuizHistory) (interface{}, error) {
	result := config.DB.Find(quizHistories)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return quizHistories, nil
}

func GetQuizHistory(quizHistory *model.QuizHistory, id string) (interface{}, error) {
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

func GetByUserIDQuizHistories(quizHistories *[]model.QuizHistory, id string) (interface{}, error) {
	idInt, errIdInt := strconv.Atoi(id)
	if errIdInt != nil {
		return nil, errors.New("id must be integer")
	}

	result := config.DB.Where("user_id = ?", idInt).Find(quizHistories)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return quizHistories, nil
}

func GetByQuizIDQuizHistories(quizHistories *[]model.QuizHistory, id string) (interface{}, error) {
	idInt, errIdInt := strconv.Atoi(id)
	if errIdInt != nil {
		return nil, errors.New("id must be integer")
	}

	result := config.DB.Where("quiz_id = ?", idInt).Find(quizHistories)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return quizHistories, nil
}

func CreateQuizHistory(quizHistory *model.QuizHistory) (interface{}, error) {
	result := config.DB.Create(quizHistory)
	err := result.Error

	if err != nil {
		return nil, err
	}

	return quizHistory, nil
}

func UpdateQuizHistory(quizHistory *model.QuizHistory, id string) (interface{}, error) {
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

func DeleteQuizHistory(id string) (interface{}, error) {
	quizHistory := model.QuizHistory{}
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

func GetSortQuizHistories(quizHistories *[]model.QuizHistory, sortBy string, order string) (interface{}, error) {
	result := config.DB.Order(sortBy + " " + order).Find(quizHistories)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}
	return quizHistories, nil
}

func GetPaginationQuizHistories(quizHistories *[]model.QuizHistory, page string, limit string) (interface{}, error) {
	pageInt, errPageInt := strconv.Atoi(page)
	if errPageInt != nil {
		return nil, errors.New("page must be integer")
	}

	limitInt, errLimitInt := strconv.Atoi(limit)
	if errLimitInt != nil {
		return nil, errors.New("limit must be integer")
	}

	offset := (pageInt - 1) * limitInt
	result := config.DB.Offset(offset).Limit(limitInt).Find(quizHistories)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}
	return quizHistories, nil
}

func GetScoreQuizHistories(quizHistories *[]model.QuizHistory, data string, score string) (interface{}, error) {
	scoreInt, errScoreInt := strconv.Atoi(score)
	if errScoreInt != nil {
		return nil, errors.New("score must be integer")
	}

	result := config.DB

	if data == "greater_than" {
		result = config.DB.Where("score > ?", scoreInt).Find(quizHistories)
	} else if data == "less_than" {
		result = config.DB.Where("score < ?", scoreInt).Find(quizHistories)
	} else if data == "greater_than_equal" {
		result = config.DB.Where("score >= ?", scoreInt).Find(quizHistories)
	} else if data == "less_than_equal" {
		result = config.DB.Where("score <= ?", scoreInt).Find(quizHistories)
	} else if data == "equal" {
		result = config.DB.Where("score = ?", scoreInt).Find(quizHistories)
	} else if data == "not_equal" {
		result = config.DB.Where("score != ?", scoreInt).Find(quizHistories)
	}

	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}
	return quizHistories, nil
}

func GetAttemptDateRangeQuizHistories(quizHistories *[]model.QuizHistory, startDate string, endDate string) (interface{}, error) {
	result := config.DB.Where("attempt_date BETWEEN ? AND ?", startDate, endDate).Find(quizHistories)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}
	return quizHistories, nil
}
