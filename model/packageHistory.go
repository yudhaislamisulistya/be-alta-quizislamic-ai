package model

import "gorm.io/gorm"

type PackageHistory struct {
	gorm.Model
	ID              uint   `json:"id" form:"id" gorm:"primary_key;auto_increment"`
	PackageID       uint   `json:"package_id" form:"package_id"`
	UserID          uint   `json:"user_id" form:"user_id"`
	TransactionDate string `json:"transaction_date" form:"transaction_date"`
	Status          string `json:"status" form:"status" gorm:"default:'pending'"`
}
