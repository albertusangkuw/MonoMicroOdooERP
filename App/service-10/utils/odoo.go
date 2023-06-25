package utils

import (
	"fmt"
	odoo "github.com/ausuwardi/godooj"
	"github.com/labstack/gommon/log"
)

var liveClient* odoo.Client

type OdooConfig struct {
	BaseURL string
	DB string
	Login string
	Password string
}
var CurrentOdooConfig OdooConfig

func InitClient()  {
	var errLogin error
	liveClient, errLogin = odoo.Connect(
		CurrentOdooConfig.BaseURL,
		CurrentOdooConfig.DB,
		CurrentOdooConfig.Login,
		CurrentOdooConfig.Password)
	if errLogin != nil {
		fmt.Println(errLogin.Error())
		fmt.Println(CurrentOdooConfig)
		log.Error("Error connecting to Odoo... !! :" +  CurrentOdooConfig.BaseURL )
	}
}

func CallOdoo() odoo.Client{
	if liveClient == nil{
		go InitClient()
		return odoo.Client{}
	}
	return *liveClient
}