package dto

type PosCategoryDTO struct {
	ID      uint    `json:"id"`
	Name  string `json:"name"`
	ParentID interface{}  `json:"parent_id"`
	ChildID  []uint   `json:"child_id"`
	Sequence int `json:"sequence"`
	Image128 interface{} `json:"image_128"`
	LastUpdate string `json:"__last_update"`
	HasImage bool `json:"has_image"`
	DisplayName string `json:"display_name"`
	CreateUID []interface{} `json:"create_uid"`
	CreateDate      string      `json:"create_date"`
	WriteUID  []interface{} `json:"write_uid"`
	WriteDate      string      `json:"write_date"`

}
