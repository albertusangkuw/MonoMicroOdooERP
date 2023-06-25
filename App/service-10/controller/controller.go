package controller

import (
	"net/http"
	"service-10/utils"

	odoo "github.com/ausuwardi/godooj"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func Index(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func DemoCallOdoo(c echo.Context) error {
	client := utils.CallOdoo()
	hasil, err := client.SearchRead(
		"product.category",
		odoo.List{},
		[]string{"id", "name"},
	)
	if err != nil {
		log.Error(err.Error())
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, hasil)

}
