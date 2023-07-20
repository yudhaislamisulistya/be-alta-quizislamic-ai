package model

import "gorm.io/gorm"

type Level struct {
	gorm.Model
	ID       int    `json:"id" form:"id" gorm:"primary_key;auto_increment"`
	Grade    string `json:"grade" form:"grade" validate:"required"`
	SubLevel string `json:"sub_level" form:"sub_level" validate:"required"`
}
