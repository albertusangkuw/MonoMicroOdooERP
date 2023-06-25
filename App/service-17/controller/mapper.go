package controller

import (
	"fmt"
	"service-17/dto"
	"service-17/model"
	"service-17/service"
)

func createMapper(data dto.CreateDTO, user model.ResUser) interface{} {
	for _, p := range data.Params.Args {
		newTag, err := service.Create(
			uint(p.Color),
			p.Name,
			user)

		if err != nil {
			fmt.Println(err.Error())
			return dto.GetResultDTO(data.ID, false)
		}
		return dto.GetResultDTO(data.ID, newTag)
	}
	return dto.GetResultDTO(data.ID, false)
}

func readMapper(data dto.ReadDTO) interface{} {
	//Bila tidak ada filter atau id batalkan
	if len(data.Params.Args) == 0 {
		return dto.GetResultDTO(data.ID, []string{})
	}
	idsArray, _ := data.Params.Args[0].([]interface{})

	uintIDs := make([]uint, len(idsArray))
	for i, id := range idsArray {
		if num, ok := id.(float64); ok {
			uintIDs[i] = uint(num)
		}
	}
	res, err := service.Read(uintIDs)
	if err != nil {
		return dto.GetResultDTO(data.ID, false)
	}
	return dto.GetResultDTO(data.ID, res)
}


func deleteMapper(data dto.ReadDTO) interface{} {
	//Bila tidak ada id batalkan
	if len(data.Params.Args) == 0 {
		return dto.GetResultDTO(data.ID, []string{})
	}
	idsArray, _ := data.Params.Args[0].([]interface{})
	for _, id := range idsArray {
		if num, ok := id.(float64); ok {
			status := service.Unlink(uint(num))
			return dto.GetResultDTO(data.ID, status)
		}
	}
	return dto.GetResultDTO(data.ID, false)
}
