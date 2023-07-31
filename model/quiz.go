package model

import "gorm.io/gorm"

type Quiz struct {
	gorm.Model
	ID          int    `json:"id" form:"id" gorm:"primary_key;auto_increment"`
	UserID      int    `json:"user_id" form:"user_id" validate:"required"`
	LevelID     int    `json:"level_id" form:"level_id" validate:"required"`
	Token       string `json:"token" form:"token" gorm:"unique"`
	Name        string `json:"name" form:"name" validate:"required"`
	Description string `json:"description" form:"description"`
	Topic       string `json:"topic" form:"topic" validate:"required"`
	Image       string `json:"image" form:"image"`
	IsDraft     bool   `json:"is_draft" form:"is_draft" gorm:"default:true"`
	IsPublished bool   `json:"is_published" form:"is_published" gorm:"default:false"`
	IsArchived  bool   `json:"is_archived" form:"is_archived" gorm:"default:false"`
	IsLocked    bool   `json:"is_locked" form:"is_locked" gorm:"default:false"`
	IsPublic    bool   `json:"is_public" form:"is_public" gorm:"default:false"`
	IsFavorite  bool   `json:"is_favorite" form:"is_favorite" gorm:"default:false"`
}
