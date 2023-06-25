package controller

import (
	"encoding/json"
	"service-10/dto"
	"service-10/model"
	"service-10/service"
)

func methodPosCategorySelector(method string, body []byte, context *dto.JWTData) interface{} {
	currModel := "pos.category"
	switch method {
	case "create":
		acc, gid := service.IsAllowedByACL(service.ModelID[currModel], context.GroupID, "c")
		if !acc {
			return map[string]string{"no access": "create"}
		}
		var createDTO dto.CreatePosCategoryDTO
		err := json.Unmarshal(body, &createDTO)
		if err != nil {
			return map[string]string{"error parse dto": err.Error()}
		}
		return createPosCategoryMapper(createDTO, model.ResUser{UID: int(context.Uid), GroupID: gid})
	case "read":
		acc, _ := service.IsAllowedByACL(service.ModelID[currModel], context.GroupID, "r")
		if !acc {
			return map[string]string{"no access": "read"}
		}
		var readDTO dto.ReadDTO
		err := json.Unmarshal(body, &readDTO)
		if err != nil {
			return map[string]string{"error parse dto": err.Error()}
		}
		return readPosCategoryMapper(readDTO)
	case "write":
		acc, gid := service.IsAllowedByACL(service.ModelID[currModel], context.GroupID, "w")
		if !acc {
			return map[string]string{"no access": "write"}
		}
		var writeDTO dto.WriteDTO
		err := json.Unmarshal(body, &writeDTO)
		if err != nil {
			return map[string]string{"error parse dto": err.Error()}
		}
		return writePosCategoryMapper(writeDTO, model.ResUser{UID: int(context.Uid), GroupID: gid})
	case "unlink":
		acc, _ := service.IsAllowedByACL(service.ModelID[currModel], context.GroupID, "u")
		if !acc {
			return map[string]string{"no access": "unlink"}
		}
		var readDTO dto.ReadDTO
		err := json.Unmarshal(body, &readDTO)
		if err != nil {
			return map[string]string{"error parse dto": err.Error()}
		}
		return deletePosCategoryMapper(readDTO)
	default:
		return nil

	}

}
