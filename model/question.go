package model

import "gorm.io/gorm"

type Questions struct {
	gorm.Model
	ID       int    `json:"id"`
	UserID   int    `json:"user_id" form:"user_id" validate:"required"`
	QuizID   int    `json:"quiz_id" form:"quiz_id" validate:"required"`
	Question string `json:"question"`
	Type     string `json:"type" form:"type" validate:"required"`
	Options  string `json:"options,omitempty" form:"options,omitempty"`
	Answer   string `json:"answer,omitempty" form:"answer,omitempty"`
	IsTrue   int    `json:"is_true"`
	Point    int    `json:"point" form:"point" validate:"required" gorm:"default:10"`
}
