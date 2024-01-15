package main

import (
	"awesomeProject/src/domains/engine"
	"fmt"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	loadEnv()
	// Listen and Server in 0.0.0.0:8080
	err := engine.MakeEngine().Run(":8080")

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
