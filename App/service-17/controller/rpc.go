package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"service-17/dto"
	"service-17/model"
	"service-17/service"

	"github.com/labstack/echo/v4"
)

var availableModel = "calendar.event.type"
var availableLang = "en_US"
var availableMethod = [...]string{"create", "write", "read", "unlink"}

func ProcessRPC(c echo.Context) error {
	var requestBody dto.CustomJSONRPCContainer

	body, err := io.ReadAll(c.Request().Body)

	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	err = json.Unmarshal(body, &requestBody)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if requestBody.Method != "call" {
		return c.JSON(http.StatusMethodNotAllowed, map[string]string{"error": "No method available for " + requestBody.Method})
	}
	if requestBody.Params.Model != availableModel {
		return c.JSON(http.StatusMethodNotAllowed, map[string]string{"error": "No model available for " + requestBody.Params.Model})
	}
	if requestBody.Params.Kwargs.Context.Lang != availableLang {
		return c.JSON(http.StatusNotImplemented, map[string]string{"error": "Current context language is not yet available for this " + requestBody.Params.Model})
	}

	for _, m := range availableMethod {
		if requestBody.Params.Method == m {
			jwtInfo, okJwt := c.Get("jwt_data").(*dto.JWTData)
			if !okJwt {
				return c.JSON(http.StatusBadRequest, map[string]string{"error": "No JWT Content "})
			}

			return c.JSON(http.StatusOK, methodSelector(m, body, jwtInfo))
		}
	}
	return c.JSON(http.StatusMethodNotAllowed, map[string]string{"error": "No method available for " + requestBody.Params.Method})

}

func methodSelector(method string, body []byte, context *dto.JWTData) interface{} {
	switch method {
	case "create":
		acc, gid := service.IsAllowedByACL(context.GroupID, "c")
		if !acc {
			return map[string]string{"no access": "create"}
		}
		var createDTO dto.CreateDTO
		err := json.Unmarshal(body, &createDTO)
		if err != nil {
			return map[string]string{"error parse dto": err.Error()}
		}
		return createMapper(createDTO, model.ResUser{UID: int(context.Uid), GroupID: gid})
	case "read":
		acc, _ := service.IsAllowedByACL(context.GroupID, "r")
		if !acc {
			return map[string]string{"no access": "read"}
		}
		var readDTO dto.ReadDTO
		err := json.Unmarshal(body, &readDTO)
		if err != nil {
			return map[string]string{"error parse dto": err.Error()}
		}
		return readMapper(readDTO)
	case "unlink":
		acc, _ := service.IsAllowedByACL(context.GroupID, "u")
		if !acc {
			return map[string]string{"no access": "unlink"}
		}
		var readDTO dto.ReadDTO
		err := json.Unmarshal(body, &readDTO)
		if err != nil {
			return map[string]string{"error parse dto": err.Error()}
		}
		return deleteMapper(readDTO)
	default:
		return nil
	}
	return nil

}
