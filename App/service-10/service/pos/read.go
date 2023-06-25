package pos

import (
	"service-10/dto"
	"service-10/model"
	"service-10/utils"
	"strconv"
)

// Get
// Browse
// Read
// SearchMode

func fullNamePosCategoryGetByID(livePosCateg map[uint]model.PosCategory, id uint) string {
	val, ok := livePosCateg[id]
	if !ok {
		return "-"
	}
	if val.ParentID == 0 {
		return val.Name
	}
	return fullNamePosCategoryGetByID(livePosCateg, val.ParentID) + " / " + val.Name
}

func updateLivePosCateg(pc []model.PosCategory) map[uint]model.PosCategory {

	livePosCateg := map[uint]model.PosCategory{}

	for _, p := range pc {
		if p.ChildID == nil {
			p.ChildID = map[uint]uint{}
		}
		_, foundOld := livePosCateg[p.ID]
		if foundOld {
			p.ChildID = livePosCateg[p.ID].ChildID
		}
		livePosCateg[p.ID] = p

		if p.ParentID != 0 {
			_, found := livePosCateg[p.ParentID]
			if !found {
				livePosCateg[p.ParentID] = model.PosCategory{
					ID:      p.ParentID,
					ChildID: map[uint]uint{},
				}
			}
			livePosCateg[p.ParentID].ChildID[p.ID] = p.ID
		}
	}
	return livePosCateg
}

func getChildAsArrayFromLive(id uint, livePosCateg map[uint]model.PosCategory) []uint {
	c := []uint{}
	for key, _ := range livePosCateg[id].ChildID {
		c = append(c, key)
	}
	return c
}

func Read(ids []uint, fields []string) ([]dto.PosCategoryDTO, error) {
	var allPosCateg []model.PosCategory
	db := utils.Database()
	db.Preload("CreateResUser").
		Preload("WriteResUser").
		Preload("Image").
		Find(&allPosCateg, ids)

	if db.Error != nil {
		return nil, db.Error
	}

	livePosCateg := updateLivePosCateg(allPosCateg)

	var allPosCategDTO = make([]dto.PosCategoryDTO, len(allPosCateg))
	for i, p := range allPosCateg {
		var pDTO = dto.PosCategoryDTO{
			ID:         p.ID,
			Name:       p.Name,
			Sequence:   p.Sequence,
			Image128:   false,
			HasImage:   p.ImageID != nil,
			CreateUID:  p.CreateResUser.ToArray(),
			CreateDate: dto.TimeToString(p.CreateAt),
			WriteUID:   p.WriteResUser.ToArray(),
			WriteDate:  dto.TimeToString(p.WriteAt),
			LastUpdate: dto.TimeToString(p.WriteAt),
		}
		if p.ParentID == 0 {
			pDTO.ParentID = false
			pDTO.DisplayName = p.DisplayName()
		} else {
			pDTO.ParentID = []interface{}{p.ParentID, livePosCateg[p.ParentID].Name}
			pDTO.DisplayName = fullNamePosCategoryGetByID(livePosCateg, p.ID)
		}

		if p.ImageID != nil {
			pDTO.Image128 = strconv.Itoa(p.Image.ImageSizeByte) + " b"
		}

		pDTO.ChildID = getChildAsArrayFromLive(p.ID, livePosCateg)

		allPosCategDTO[i] = pDTO
	}

	return allPosCategDTO, nil
}
