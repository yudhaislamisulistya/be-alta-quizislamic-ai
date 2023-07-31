package model

import "gorm.io/gorm"

type Wallet struct {
	gorm.Model
	ID      uint   `json:"id" form:"id" gorm:"primary_key;auto_increment"`
	UserID  uint   `json:"user_id" form:"user_id" gorm:"primary_key"`
	Balance int64  `json:"balance" form:"balance" gorm:"default:50"`
	Token   string `json:"token" form:"token" gorm:"default:null"`
}
