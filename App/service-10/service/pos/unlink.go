package pos

import (
	"fmt"
	"service-10/model"
	"service-10/utils"
)

func Unlink(id uint) bool {
	db := utils.Database()
	target := model.PosCategory{ID: id}
	result := db.Preload("Image").First(&target)
	if result.RowsAffected == 0 {
		return true
	}
	if target.Image.ID != 0 && target.Image.AttachmentId != 0 {
		DeleteImageRemote(target.Image.AttachmentId)
	}
	db.Delete(&target)
	return !IsExist(target)
}

func DeleteImageRemote(attachmentID int) bool {
	client := utils.CallOdoo()
	if !client.IsValid() {
		fmt.Println("no Odoo Server to Delete Image")
		return false
	}
	_, err := client.Delete("ir.attachment", []int{attachmentID})
	if err != nil {
		fmt.Println("Image Failed to Deleted")
		return false
	}
	return true
}
