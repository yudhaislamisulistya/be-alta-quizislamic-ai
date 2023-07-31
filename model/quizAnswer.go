package model

import "gorm.io/gorm"

type QuizAnswer struct {
	gorm.Model
	ID            uint `json:"id" form:"id" gorm:"primary_key;auto_increment"`
	QuizHistoryID uint `json:"quiz_history_id" form:"quiz_history_id" gorm:"primary_key"`
	QuestionID    uint `json:"question_id" form:"question_id" gorm:"primary_key"`
	IsCorrect     bool `json:"is_correct" form:"is_correct" gorm:"default:false"`
}
