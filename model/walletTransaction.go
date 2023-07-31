package model

import "gorm.io/gorm"

type WalletTransaction struct {
	gorm.Model
	ID              uint   `json:"id" form:"id" gorm:"primary_key;auto_increment"`
	WalletID        uint   `json:"wallet_id" form:"wallet_id"`
	Amount          int64  `json:"amount" form:"amount"`
	TransactionType string `json:"transaction_type" form:"transaction_type" gorm:"default:'Top Up'"`
	TransactionDate string `json:"transaction_date" form:"transaction_date"`
}
