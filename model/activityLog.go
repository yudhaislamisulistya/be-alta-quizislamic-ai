package model

import "gorm.io/gorm"

type ActivityLog struct {
	gorm.Model
	ID               uint   `json:"id" form:"id" gorm:"primary_key;auto_increment"`
	UserID           uint   `json:"user_id" form:"user_id"`
	ActivityType     string `json:"activity_type" form:"activity_type"`
	ActivityDateTime string `json:"activity_date_time" form:"activity_date_time"`
}
