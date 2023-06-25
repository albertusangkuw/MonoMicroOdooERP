package main

import (
	"flag"
	"fmt"
	"net/http"
	"service-17/controller"
	"service-17/model"
	"service-17/service"
	"service-17/utils"

	"github.com/labstack/gommon/log"
	"golang.org/x/net/http2"

	"time"

	"github.com/go-co-op/gocron"
	"github.com/labstack/echo/v4"
)

func main() {
	// Define command-line flags
	targetEnvFile := flag.String("env", "", "Target ENV")
	flag.Parse()
	if *targetEnvFile != "" {
		fmt.Println("Target .env: " + *targetEnvFile)
		utils.LoadEnv(*targetEnvFile)
	} else {
		fmt.Println("Default .env")
		utils.LoadEnv(".env")
	}

	utils.CurrentOdooConfig = utils.OdooConfig{
		BaseURL:  utils.EnvVar("ODOO_HOST"),
		DB:       utils.EnvVar("ODOO_DB"),
		Login:    utils.EnvVar("ODOO_LOGIN"),
		Password: utils.EnvVar("ODOO_PASSWORD"),
	}
	utils.CurrentDBConfig = utils.DBConfig{
		Host:     utils.EnvVar("MY_SQL_HOST"),
		DB:       utils.EnvVar("MY_SQL_DATABASE"),
		Username: utils.EnvVar("MY_SQL_USERNAME"),
		Password: utils.EnvVar("MY_SQL_PASSWORD"),
	}

	utils.Database().AutoMigrate(&model.ResUser{}, &model.MeetingType{})

	e := echo.New()
	e.GET("/", controller.Index)
	e.POST("/jsonrpc/*", controller.ProcessRPC)
	e.Use(controller.JWTHandler)
	e.Use(controller.LoggerHandler)

	//Start Odoo Client
	utils.InitClient()
	service.SearchModelID(utils.EnvVar("MODEL_NAME"))

	backgroundTask := gocron.NewScheduler(time.UTC)
	backgroundTask.Every(1).Minute().Do(service.UpdateACL)
	backgroundTask.StartAsync()

	s := &http2.Server{}
	if err := e.StartH2CServer(":"+utils.EnvVar("PORT_APP"), s); err != http.ErrServerClosed {
		log.Fatal(err)
	}

}
