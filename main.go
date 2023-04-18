package main

import (
	"fmt"
	"log"
	"os"

	"github.com/volyanyk/todoist"

	"github.com/joho/godotenv"
)

func init() {
	// loads values from .env into the system
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	apiToken, exists := os.LookupEnv("APITOKEN")
	if !exists {
		panic("Не указан API token todoist'a")
	}

	api := todoist.New(apiToken)
	projects, err := api.GetProjects()
	if err != nil {
		return
	}

	fmt.Printf("%v", projects)
}
