package controller

import (
	"encoding/json"
	"io"
	"net/http"
	"service-10/dto"

	"github.com/labstack/echo/v4"
)

var AvailableModel = map[string]int{
	"pos.category": 1,
	"product.tag":  2,
}
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
	_, modelExist := AvailableModel[requestBody.Params.Model]
	if !modelExist {
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

			return c.JSON(http.StatusOK, modelSelector(requestBody.Params.Model, m, body, jwtInfo))
		}
	}
	return c.JSON(http.StatusMethodNotAllowed, map[string]string{"error": "No method available for " + requestBody.Params.Method})

}

func modelSelector(model string, method string, body []byte, context *dto.JWTData) interface{} {

	switch model {
	case "pos.category":
		return methodPosCategorySelector(method, body, context)
	case "product.tag":
		return methodProductTagSelector(method, body, context)
	default:
		return map[string]string{"Error No Method Selector": model}
	}
}
