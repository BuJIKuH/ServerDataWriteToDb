package main

import (
	"awesomeProject/config"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {

	errDb := config.ConnectDataBase()
	if errDb != nil {
		log.Fatalf("No connect", errDb)
	}
	config.ResponseServer()
}
