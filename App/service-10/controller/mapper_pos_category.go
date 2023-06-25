package controller

import (
	"fmt"
	"service-10/dto"
	"service-10/model"
	"service-10/service/pos"

	odoo "github.com/ausuwardi/godooj"
)

// Format Data Interface "args": [
//
//	    {
//	        "image_128": "false/base64",
//	        "name": "PosCategoryTest",
//	        "parent_id": 3
//	    }
//	],
func createPosCategoryMapper(data dto.CreatePosCategoryDTO, user model.ResUser) interface{} {
	for _, p := range data.Params.Args {
		parentID, _ := odoo.Val2Int(p.ParentID)
		image128, _ := odoo.Val2String(p.Image128)
		newPos, err := pos.Create(
			p.Name,
			uint(parentID),
			p.Sequence,
			image128,
			user)

		if err != nil {
			fmt.Println(err.Error())
			return dto.GetResultDTO(data.ID, false)
		}
		//Contoh return
		//{
		//	"jsonrpc": "2.0",
		//	"id": 26,
		//	"result": 17
		//}
		return dto.GetResultDTO(data.ID, newPos)
	}
	return dto.GetResultDTO(data.ID, false)
}

// Format Data Interface  "args": [
//
//	    [
//	        3
//	    ],
//	    [
//	        "__last_update",
//	        "image_128",
//	        "name",
//	        "parent_id",
//	        "display_name"
//	    ]
//	],
func readPosCategoryMapper(data dto.ReadDTO) interface{} {
	//Bila tidak ada filter atau id batalkan
	if len(data.Params.Args) == 0 {
		return dto.GetResultDTO(data.ID, []string{})
	}
	idsArray, _ := data.Params.Args[0].([]interface{})

	var columns []string
	//Column filter return
	if len(data.Params.Args) == 2 {
		columns, _ = data.Params.Args[1].([]string)
	}

	uintIDs := make([]uint, len(idsArray))
	for i, id := range idsArray {
		if num, ok := id.(float64); ok {
			uintIDs[i] = uint(num)
		}
	}
	res, err := pos.Read(uintIDs, columns)
	if err != nil {
		return dto.GetResultDTO(data.ID, false)
	}
	//Contoh return
	//{
	//	"jsonrpc": "2.0",
	//	"id": 26,
	//	"result": [
	//        {
	//            "id": 3,
	//            "__last_update": "2023-05-18 12:55:48",
	//            "image_128": false,
	//            "name": "Chairs",
	//            "parent_id": false,
	//            "display_name": "Chairs"
	//        }
	//    ]
	//}
	return dto.GetResultDTO(data.ID, res)
}

func writePosCategoryMapper(data dto.WriteDTO, rUser model.ResUser) interface{} {
	idData, posUpdate, targetUpdate := data.GetData()
	if idData == 0 {
		return dto.GetResultDTO(data.ID, false)
	}
	if len(targetUpdate) == 0 {
		return dto.GetResultDTO(data.ID, false)
	}
	err := pos.Write(idData, posUpdate, targetUpdate, rUser)
	if err != nil {
		return dto.GetResultDTO(data.ID, false)
	}
	return dto.GetResultDTO(data.ID, true)
}

func deletePosCategoryMapper(data dto.ReadDTO) interface{} {
	//Bila tidak ada id batalkan
	if len(data.Params.Args) == 0 {
		return dto.GetResultDTO(data.ID, []string{})
	}
	idsArray, _ := data.Params.Args[0].([]interface{})
	for _, id := range idsArray {
		if num, ok := id.(float64); ok {
			status := pos.Unlink(uint(num))
			return dto.GetResultDTO(data.ID, status)
		}
	}
	return dto.GetResultDTO(data.ID, false)
}
