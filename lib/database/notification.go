package database

import (
	"errors"
	"project/config"
	"project/model"
	"strconv"
)

func GetNotifications(notifications *[]model.Notification) (interface{}, error) {
	result := config.DB.Find(&notifications)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return notifications, nil
}

func GetNotification(Notification *model.Notification, id string) (interface{}, error) {
	idInt, errIdInt := strconv.Atoi(id)
	if errIdInt != nil {
		return nil, errors.New("id must be integer")
	}

	result := config.DB.First(Notification, idInt)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return Notification, nil
}

func GetByUserIDNotifications(notifications *[]model.Notification, id string) (interface{}, error) {
	idInt, errIdInt := strconv.Atoi(id)
	if errIdInt != nil {
		return nil, errors.New("id must be integer")
	}

	result := config.DB.Where("user_id = ?", idInt).Find(notifications)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return notifications, nil
}

func CreateNotification(Notification *model.Notification) (interface{}, error) {
	result := config.DB.Create(Notification)
	err := result.Error

	if err != nil {
		return nil, err
	}

	return Notification, nil
}

func UpdateNotification(Notification *model.Notification, id string) (interface{}, error) {
	idInt, errIdInt := strconv.Atoi(id)
	if errIdInt != nil {
		return nil, errors.New("id must be integer")
	}

	result := config.DB.Model(Notification).Where("id = ?", idInt).Updates(Notification)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return Notification, nil
}

func DeleteNotification(id string) (interface{}, error) {
	Notification := model.Notification{}
	idInt, errIdInt := strconv.Atoi(id)
	if errIdInt != nil {
		return nil, errors.New("id must be integer")
	}

	result := config.DB.Delete(&Notification, idInt)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return len, nil
}

func GetSortNotifications(notifications *[]model.Notification, sortBy string, order string) (interface{}, error) {
	result := config.DB.Order(sortBy + " " + order).Find(notifications)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}
	return notifications, nil
}

func GetPaginationNotifications(notifications *[]model.Notification, page string, limit string) (interface{}, error) {
	pageInt, errPageInt := strconv.Atoi(page)
	if errPageInt != nil {
		return nil, errors.New("page must be integer")
	}

	limitInt, errLimitInt := strconv.Atoi(limit)
	if errLimitInt != nil {
		return nil, errors.New("limit must be integer")
	}

	offset := (pageInt - 1) * limitInt
	result := config.DB.Offset(offset).Limit(limitInt).Find(notifications)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}
	return notifications, nil
}

func GetFilterNotifications(notifications *[]model.Notification, activityType string) (interface{}, error) {
	result := config.DB.Where("notification_type = ?", activityType).Find(notifications)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return notifications, nil
}
