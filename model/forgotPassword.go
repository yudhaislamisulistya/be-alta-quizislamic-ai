package model

import "gorm.io/gorm"

type ForgotPassword struct {
	gorm.Model
	ID                     uint   `json:"id" form:"id" gorm:"primary_key;auto_increment"`
	UserID                 uint   `json:"user_id" form:"user_id"`
	VerificationCode       string `json:"verification_code" form:"verification_code"`
	IsUsedVerificationCode bool   `json:"is_used_verification_code" form:"is_used_verification_code" gorm:"default:false"`
	ExpiredAtVerification  int    `json:"expired_at_verification" form:"expired_at_verification"`
}
