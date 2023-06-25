package pos

import (
	"errors"
	"fmt"
	"service-10/dto"
	"service-10/model"
	"service-10/utils"
	"time"

	odoo "github.com/ausuwardi/godooj"
)

func IsExist(category model.PosCategory) bool {
	result := utils.Database().First(&category)
	if result.RowsAffected == 0 {
		return false
	}
	return true
}

func Write(id uint, p dto.PosCategoryRequest, targetUpdate []string, uid model.ResUser) error {
	posCateg := model.PosCategory{ID: id}

	pos := model.PosCategory{
		ID:           id,
		Name:         p.Name,
		Sequence:     p.Sequence,
		WriteResUser: uid,
		WriteAt:      time.Now(),
	}

	result := utils.Database().First(&posCateg)
	if result.RowsAffected == 0 {
		return errors.New("Not Found")
	}

	currParentID, errParentID := odoo.Val2Int(p.ParentID)
	if !errParentID {
		parent := model.PosCategory{ID: uint(currParentID)}
		if !IsExist(parent) {
			return errors.New("Parent Not Found")
		}
		pos.ParentID = uint(currParentID)
		if isRecursion(pos, parent) {
			return errors.New("Cyclic Dependency Detected ")
		}
	}

	tx := utils.Database().Begin()
	CreateUserIfNotFound(tx, uid)
	pos.WriteUID = &uid.UID

	targetUpdate = append(targetUpdate, "write_uid", "write_at")
	image128, _ := odoo.Val2String(p.Image128)
	if image128 == "" {
		pos.ImageID = nil
		res := tx.Model(&pos).Select(targetUpdate).Updates(pos)
		if res.Error != nil {
			tx.Rollback()
			return res.Error
		}
		tx.Commit()
		return nil
	}

	newImg := model.PosCategoryImage{
		ImageBase64:   image128,
		ImageSizeByte: len(image128),
	}
	resImg := tx.Create(&newImg)
	if resImg.Error != nil {
		tx.Rollback()
		fmt.Println(resImg.Error.Error())
		return resImg.Error
	}

	pos.Image = newImg
	res := tx.Create(&pos)
	if res.Error != nil {
		tx.Rollback()
		return res.Error
	}
	tx.Commit()
	go SaveImageRemote(newImg, int(pos.ID))
	return nil
}
