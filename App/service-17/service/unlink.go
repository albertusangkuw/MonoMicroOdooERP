package service

import (
	"service-17/model"
	"service-17/utils"
)

func Unlink(id uint) bool {
	db := utils.Database()
	target := model.MeetingType{ID: id}
	result := db.First(&target)
	if result.RowsAffected == 0 {
		return true
	}
	db.Delete(&target)
	return !IsExist(target)
}
