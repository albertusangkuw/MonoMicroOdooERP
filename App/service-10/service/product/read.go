package product

import (
	"service-10/dto"
	"service-10/model"
	"service-10/utils"

	odoo "github.com/ausuwardi/godooj"
	"github.com/labstack/gommon/log"
)

func getProductTemplate(id uint) []int {
	od := utils.CallOdoo()
	if !od.IsValid() {
		return []int{}
	}
	records, err := od.SearchRead("product.template", odoo.List{
		odoo.List{"product_tag_ids", "=", id},
	},
		[]string{"id", "product_variant_ids"})
	if err != nil {
		log.Error("Error retrieving data")
		return []int{}
	}
	productID := make([]int, len(records))
	for i, p := range records {
		productID[i], _ = odoo.IntField(p, "id")
	}
	return productID
}

func getProductProduct(id uint) []int {
	od := utils.CallOdoo()
	if !od.IsValid() {
		return []int{}
	}
	records, err := od.SearchRead("product.product", odoo.List{
		odoo.List{"all_product_tag_ids", "=", id},
	},
		[]string{"id"})
	if err != nil {
		log.Error("Error retrieving data")
		return []int{}
	}
	productID := make([]int, len(records))
	for i, p := range records {
		productID[i], _ = odoo.IntField(p, "id")
	}
	return productID
}

func computeProductIds(pt []int, pp []int) []int {
	c := []int{}
	for _, v := range pt {
		found := false
		for _, t := range pp {
			if t == v {
				found = true
				break
			}
		}
		if found {
			c = append(c, v)
		}
	}
	return c
}

func Read(ids []uint) ([]dto.ProductTagDTO, error) {
	var tags []model.ProductTag
	db := utils.Database()
	db.Preload("CreateResUser").
		Preload("WriteResUser").
		Find(&tags, ids)

	if db.Error != nil {
		return nil, db.Error
	}
	var tagDTO = make([]dto.ProductTagDTO, len(tags))
	for i, p := range tags {
		var pDTO = dto.ProductTagDTO{
			ID:          p.ID,
			Name:        p.Name,
			Color:       p.Color,
			DisplayName: p.Name,
			CreateUID:   p.CreateResUser.ToArray(),
			CreateDate:  dto.TimeToString(p.CreateAt),
			WriteUID:    p.WriteResUser.ToArray(),
			WriteDate:   dto.TimeToString(p.WriteAt),
			LastUpdate:  dto.TimeToString(p.WriteAt),
		}
		pDTO.ProductTemplateIds = getProductTemplate(p.ID)
		pDTO.ProductIds = getProductProduct(p.ID)
		pDTO.ProductProductIds = computeProductIds(pDTO.ProductTemplateIds, pDTO.ProductIds)
		tagDTO[i] = pDTO
	}
	return tagDTO, nil
}
