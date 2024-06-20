package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

var DbUser string
var DbPassword string
var DbHost string
var DbPort string
var DbName string
var ServerPort string

func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	DbUser = os.Getenv("DB_USER")
	DbPassword = os.Getenv("DB_PASSWORD")
	DbHost = os.Getenv("DB_HOST")
	DbPort = os.Getenv("DB_PORT")
	DbName = os.Getenv("DB_NAME")
	ServerPort = os.Getenv("SERVER_PORT")

	log.Println("Config initialized")
}
