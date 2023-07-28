package model

import "gorm.io/gorm"

type Notification struct {
	gorm.Model
	ID               uint   `json:"id" form:"id" gorm:"primary_key;auto_increment"`
	UserID           uint   `json:"user_id" form:"user_id"`
	NotificationType string `json:"notification_type" form:"notification_type"`
	NotificationText string `json:"notification_text" form:"notification_text"`
	IsRead           bool   `json:"is_read" form:"is_read"`
}
