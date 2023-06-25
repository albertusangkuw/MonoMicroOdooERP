package pos

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

func Create(name string, parentID uint, sequence int, imageBase64 string, uid model.ResUser) (uint, error) {
	pos := model.PosCategory{
		Name:          name,
		Sequence:      sequence,
		CreateAt:      time.Now(),
		CreateResUser: uid,
		WriteResUser:  uid,
		WriteAt:       time.Now(),
	}
	var parentRef model.PosCategory
	if parentID != 0 {
		//Get parent
		parentRef.ID = parentID
		result := utils.Database().First(&parentRef, parentID)
		if result.Error != nil {
			//No Parent
			return 0, result.Error
		}
		pos.ParentID = parentRef.ID
	}

	tx := utils.Database().Begin()
	CreateUserIfNotFound(tx, uid)

	if imageBase64 == "" {
		pos.ImageID = nil
		res := tx.Create(&pos)
		if res.Error != nil {
			tx.Rollback()
			return 0, res.Error
		}
		tx.Commit()
		return pos.ID, nil
	}

	newImg := model.PosCategoryImage{
		ImageBase64:   imageBase64,
		ImageSizeByte: len(imageBase64),
	}
	resImg := tx.Create(&newImg)
	if resImg.Error != nil {
		tx.Rollback()
		fmt.Println(resImg.Error.Error())
		return 0, resImg.Error
	}
	pos.Image = newImg
	res := tx.Create(&pos)
	if res.Error != nil {
		tx.Rollback()
		return 0, res.Error
	}
	tx.Commit()
	//Background Task mengirimkan data
	go SaveImageRemote(newImg, int(pos.ID))
	return pos.ID, nil
}

func SaveImageRemote(image model.PosCategoryImage, posID int) {
	client := utils.CallOdoo()
	if !client.IsValid() {
		fmt.Println("no Odoo Server to Save Image")
		return
	}
	resRPC, err := client.Create(
		"ir.attachment",
		map[string]interface{}{
			"name":        "image_128",
			"datas":       image.ImageBase64,
			"store_fname": "file_name.png",
			"res_model":   "pos.category",
			"res_field":   "image_128",
			"res_id":      posID,
		})
	if err != nil {
		fmt.Println("Image Failed to Save")
		return
	}
	// Update di DB tag Attachment ID
	db := utils.Database()
	db.Model(&image).Update("AttachmentId", resRPC)
}

func traverseNode(results *[]model.PosCategory, current uint, target uint) bool {
	//Ketika node menjadi parent maka sudah tidak valid
	if current == target {
		return true
	}
	for _, row := range *results {
		if row.ID == current {
			if row.ParentID == 0 {
				return false
			}
			return traverseNode(results, row.ParentID, target)
		}
	}
	return false
}

// p->c	     3->4
// 1->2,2->3,3->null,4->1
// 1, 3
// 2, 3
// 3, 3

func isRecursion(c model.PosCategory, p model.PosCategory) bool {
	db := utils.Database()
	if c.ParentID == 0 || p.ParentID == 0 {
		return false
	}
	var results []model.PosCategory
	res := db.Model(model.PosCategory{}).Select("id, parentID").Find(&results)
	if res.Error != nil {
		// Tidak dapat diverifikasi
		return true
	}
	return traverseNode(&results, p.ParentID, c.ID)
}
