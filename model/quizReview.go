package model

import "gorm.io/gorm"

type QuizReview struct {
	gorm.Model
	ID         uint   `json:"id" form:"id" gorm:"primary_key;auto_increment"`
	UserID     uint   `json:"user_id" form:"user_id"`
	QuizID     uint   `json:"quiz_id" form:"quiz_id"`
	ReviewText string `json:"review_text" form:"review_text"`
	Rating     int64  `json:"rating" form:"rating"`
}
