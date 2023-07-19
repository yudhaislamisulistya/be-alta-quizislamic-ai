package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID                 uint      `json:"id" form:"id" gorm:"primary_key;auto_increment"`
	UUID               uuid.UUID `json:"uuid" form:"uuid" gorm:"primary_key"`
	Username           string    `json:"username" form:"username" gorm:"unique" validate:"required,min=5,max=20"`
	Email              string    `json:"email" form:"email" gorm:"unique" validate:"required,email"`
	Password           string    `json:"password" form:"password" validate:"required,min=8,max=20"`
	Name               string    `json:"name" form:"name" validate:"required"`
	Address            string    `json:"address" form:"address"`
	BirthDate          string    `json:"birth_date" form:"birth_date"`
	Gender             string    `json:"gender" form:"gender"`
	PhoneNumber        string    `json:"phone_number" form:"phone_number"`
	ProfilePhoto       string    `json:"profile_photo" form:"profile_photo"`
	JoinedAt           string    `json:"joined_at" form:"joined_at"`
	LastLogin          string    `json:"last_login" form:"last_login"`
	IsAdmin            bool      `json:"is_admin" form:"is_admin" gorm:"default:false"`
	AccountStatus      string    `json:"account_status" form:"account_status" gorm:"default:active"`
	PIN                string    `json:"pin" form:"pin" gorm:"default:000000"`
	RegisteredVia      string    `json:"registered_via" form:"registered_via" gorm:"default:email"`
	Token              string    `json:"token" form:"token"`
	TokenExpired       int64     `json:"token_expired" form:"token_expired" gorm:"default:null"`
	TokenVerifiedEmail string    `json:"token_verified_email" form:"token_verified_email"`
	IsVerifiedEmail    bool      `json:"is_verified_email" form:"is_verified_email" gorm:"default:false"`
}
