package product

import (
	"service-10/model"
	"service-10/utils"
)

func Unlink(id uint) bool {
	db := utils.Database()
	target := model.ProductTag{ID: id}
	result := db.First(&target)
	if result.RowsAffected == 0 {
		return true
	}
	db.Delete(&target)
	return !IsExist(target)
}
