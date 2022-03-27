package main

import (
	"chk/controller"
	database "chk/infrastructure"
	"fmt"
	"github.com/getsentry/sentry-go"
	"github.com/gin-gonic/gin"
	_ "github.com/joho/godotenv/autoload"
	"os"
)

var Mode = os.Getenv("MODE")
var port = os.Getenv("PORT")
var dbHost = os.Getenv("DB_HOST")
var dbPort = os.Getenv("DB_PORT")
var dbDatabase = os.Getenv("DB_DATABASE")
var dbUsername = os.Getenv("DB_USERNAME")
var dbPassword = os.Getenv("DB_PASSWORD")
var SentryLink = ""

func init() {

	gin.SetMode(Mode)

	if err := sentry.Init(sentry.ClientOptions{
		Dsn: SentryLink,
	}); err != nil {

	}
	dbInfo := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", dbUsername, dbPassword, dbHost, dbPort, dbDatabase)
	database.Init(dbInfo)
}

func main() {
	router := gin.Default()
	homeController := new(controller.HomeController)

	router.POST("/data", homeController.Upload)
	router.GET("/data", homeController.Get)

	router.Run(":" + port)
}
