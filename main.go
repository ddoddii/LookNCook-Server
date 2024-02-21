package main

import (
	"github.com/gdsc-ys/lookncook-server/config"
	"github.com/gdsc-ys/lookncook-server/src/router"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("unable to locate .env file")
	}
}

func main() {
	init := config.InitializeApp()
	echoR := router.Router(init)

	echoR.Logger.Fatal(echoR.Start(":8080"))
}
