package dto

type ProductTagRequest struct {
	Name              string `json:"name"`
	Color             int `json:"color"`
}

type CreateProductTagDTO struct{
	ID      int    `json:"id"`
	Params  struct {
		Args   []ProductTagRequest `json:"args"`
		Kwargs struct {
			Context interface{}`json:"context"`
			Domain interface{}`json:"domain"`
			Field interface{}`json:"domain"`
		} `json:"kwargs"`
	} `json:"params"`
}

