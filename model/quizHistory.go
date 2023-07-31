package model

import "gorm.io/gorm"

type QuizHistory struct {
	gorm.Model
	ID          uint   `json:"id" form:"id" gorm:"primary_key;auto_increment"`
	UserID      uint   `json:"user_id" form:"user_id" gorm:"primary_key"`
	QuizID      uint   `json:"quiz_id" form:"quiz_id" gorm:"primary_key"`
	AttemptDate string `json:"attempt_date" form:"attempt_date"`
	Score       int64  `json:"score" form:"score" gorm:"default:0"`
	IsFinished  bool   `json:"is_finished" form:"is_finished" gorm:"default:false"`
}
