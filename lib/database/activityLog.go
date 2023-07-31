package database

import (
	"errors"
	"project/config"
	"project/model"
	"strconv"
)

func GetActivityLogs(activityLogs *[]model.ActivityLog) (interface{}, error) {
	result := config.DB.Find(&activityLogs)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return activityLogs, nil
}

func GetActivityLog(ActivityLog *model.ActivityLog, id string) (interface{}, error) {
	idInt, errIdInt := strconv.Atoi(id)
	if errIdInt != nil {
		return nil, errors.New("id must be integer")
	}

	result := config.DB.First(ActivityLog, idInt)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return ActivityLog, nil
}

func GetByUserIDActivityLogs(activityLogs *[]model.ActivityLog, id string) (interface{}, error) {
	idInt, errIdInt := strconv.Atoi(id)
	if errIdInt != nil {
		return nil, errors.New("id must be integer")
	}

	result := config.DB.Where("user_id = ?", idInt).Find(activityLogs)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return activityLogs, nil
}

func CreateActivityLog(ActivityLog *model.ActivityLog) (interface{}, error) {
	result := config.DB.Create(ActivityLog)
	err := result.Error

	if err != nil {
		return nil, err
	}

	return ActivityLog, nil
}

func UpdateActivityLog(ActivityLog *model.ActivityLog, id string) (interface{}, error) {
	idInt, errIdInt := strconv.Atoi(id)
	if errIdInt != nil {
		return nil, errors.New("id must be integer")
	}

	result := config.DB.Model(ActivityLog).Where("id = ?", idInt).Updates(ActivityLog)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return ActivityLog, nil
}

func DeleteActivityLog(id string) (interface{}, error) {
	ActivityLog := model.ActivityLog{}
	idInt, errIdInt := strconv.Atoi(id)
	if errIdInt != nil {
		return nil, errors.New("id must be integer")
	}

	result := config.DB.Delete(&ActivityLog, idInt)
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

func GetSortActivityLogs(activityLogs *[]model.ActivityLog, sortBy string, order string) (interface{}, error) {
	result := config.DB.Order(sortBy + " " + order).Find(activityLogs)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}
	return activityLogs, nil
}

func GetPaginationActivityLogs(activityLogs *[]model.ActivityLog, page string, limit string) (interface{}, error) {
	pageInt, errPageInt := strconv.Atoi(page)
	if errPageInt != nil {
		return nil, errors.New("page must be integer")
	}

	limitInt, errLimitInt := strconv.Atoi(limit)
	if errLimitInt != nil {
		return nil, errors.New("limit must be integer")
	}

	offset := (pageInt - 1) * limitInt
	result := config.DB.Offset(offset).Limit(limitInt).Find(activityLogs)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}
	return activityLogs, nil
}

func GetFilterActivityLogs(activityLogs *[]model.ActivityLog, activityType string) (interface{}, error) {
	result := config.DB.Where("activity_type = ?", activityType).Find(activityLogs)
	err := result.Error
	len := result.RowsAffected

	if err != nil {
		return nil, err
	}

	if len == 0 {
		return len, nil
	}

	return activityLogs, nil
}
