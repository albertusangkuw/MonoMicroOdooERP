package dto

type ProductTagDTO struct {
	ID      uint    `json:"id"`
	Name  string `json:"name"`
	Color  uint `json:"color"`
	ProductTemplateIds []int  `json:"product_template_ids"`
	ProductProductIds []int  `json:"product_product_ids"`
	ProductIds []int  `json:"product_ids"`

	DisplayName string `json:"display_name"`
	CreateUID []interface{} `json:"create_uid"`
	CreateDate      string      `json:"create_date"`
	WriteUID  []interface{} `json:"write_uid"`
	WriteDate      string      `json:"write_date"`
	LastUpdate string `json:"__last_update"`
}


