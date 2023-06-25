package dto

import odoo "github.com/ausuwardi/godooj"


type PosCategoryRequest struct {
	Name                   string `json:"name"`
	Image128               interface{} `json:"image_128"`
	ParentID      			interface{} `json:"parent_id"`
	Sequence      			int `json:"sequence,omitempty"`
}


type CreatePosCategoryDTO struct{
	ID      int    `json:"id"`
	Params  struct {
		Args   []PosCategoryRequest `json:"args"`
		Kwargs struct {
			Context interface{}`json:"context"`
			Domain interface{}`json:"domain"`
			Field interface{}`json:"domain"`
		} `json:"kwargs"`
	} `json:"params"`
}


func (w *WriteDTO)GetData() (uint, PosCategoryRequest, []string){
	if len(w.Params.Args) != 2{
		return 0, PosCategoryRequest{}, []string{}
	}
	arg1 := w.Params.Args[0].([]interface{})
	arg2 := w.Params.Args[1].(interface{})

	var idData uint
	for _, item := range arg1 {
		idData = uint(item.(float64))
		break
	}
	targetUpdate := []string{}
	posUpdate := PosCategoryRequest{}
	name,  eName := odoo.StringField( arg2, "name")
	if eName == nil {
		targetUpdate = append(targetUpdate, "name")
		posUpdate.Name = name
	}
	sequence,  eSeq := odoo.FloatField( arg2, "sequence")
	if eSeq == nil {
		targetUpdate = append(targetUpdate, "sequence")
		posUpdate.Sequence = int(sequence)
	}

	parent,  ePar := odoo.Many2oneField(arg2, "parent_id")
	if ePar == nil {
		targetUpdate = append(targetUpdate, "sequence")
		posUpdate.ParentID = parent.ID
	}

    img, eImg := odoo.StringField(arg2,"image_128")
	if eImg == nil {
		targetUpdate = append(targetUpdate, "image_id")
		posUpdate.Image128 = img
	}
	return idData,posUpdate, targetUpdate
}
