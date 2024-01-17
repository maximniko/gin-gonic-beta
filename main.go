package main

import (
	"awesomeProject/src/domains/app/components"
	"awesomeProject/src/domains/app/models"
	"awesomeProject/src/domains/engine"
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	loadEnv()
	// Listen and Server in 0.0.0.0:8080
	err := engine.MakeEngine(makeApp()).Run(":8080")

	if err != nil {
		fmt.Println(err)
		return
	}
}

func loadEnv() {
	err := godotenv.Load(".env.local")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func makeApp() *models.App {
	app := models.NewApp()
	app.SetDbRead(components.NewMysqlConnectRead())
	app.SetGoqu(components.NewGoqu(app))
	return app
}
