package main

import (
	"log"
	"todoistapi/internal/asana"

	"github.com/joho/godotenv"
)

func init() {
	// Загружаем в систему переменные из .env
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	asanaClient := asana.NewClient()

	asanaUserId, err := asana.GetUserIdByName(asanaClient)
	if err != nil {
		panic(err)
	}

	asana.GetTaskByUserId(asanaClient, asanaUserId)
}
