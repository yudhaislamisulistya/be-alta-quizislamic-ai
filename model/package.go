package model

import "gorm.io/gorm"

type Package struct {
	gorm.Model
	ID            int    `json:"id" form:"id" gorm:"primary_key;auto_increment"`
	Name          string `json:"name" form:"name" validate:"required"`
	Type          string `json:"type" form:"type" validate:"required"`
	BalanceAmount int64  `json:"balance_amount" form:"balance_amount" validate:"required"`
	Price         int64  `json:"price" form:"price" validate:"required"`
	Description   string `json:"description" form:"description" validate:"required"`
}
