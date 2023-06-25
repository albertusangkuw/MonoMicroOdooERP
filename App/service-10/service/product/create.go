package product

import (
	"fmt"
	"service-10/model"
	"service-10/utils"
	"time"

	"gorm.io/gorm"
)

func CreateUserIfNotFound(db *gorm.DB, res_user model.ResUser) model.ResUser {
	db.FirstOrCreate(&res_user, res_user)
	return res_user
}

func Create(color uint, name string, uid model.ResUser) (uint, error) {
	tag := model.ProductTag{
		Name:          name,
		Color:         color,
		CreateAt:      time.Now(),
		CreateResUser: uid,
		WriteResUser:  uid,
		WriteAt:       time.Now(),
	}
	tx := utils.Database().Begin()
	CreateUserIfNotFound(tx, uid)

	resTag := tx.Create(&tag)
	if resTag.Error != nil {
		tx.Rollback()
		fmt.Println(resTag.Error.Error())
		return 0, resTag.Error
	}
	tx.Commit()
	return tag.ID, nil
}
