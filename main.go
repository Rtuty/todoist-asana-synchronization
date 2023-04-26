package main

import (
	"log"
	"todoistapi/internal/asana"

	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	asana.GetAllUsers()
}
