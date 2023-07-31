package model

import "gorm.io/gorm"

type QuestionCategory struct {
	gorm.Model
	ID   int    `json:"id"`
	Name string `json:"name" validate:"required"`
}
