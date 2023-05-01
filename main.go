package main

import (
	"fmt"
	"log"
	"todoistapi/internal/todoist"

	"github.com/joho/godotenv"
)

func init() {
	// Загружаем в систему переменные из .env
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func main() {
	tdistClient, err := todoist.NewClient()
	if err != nil {
		panic("new client error")
	}

	tdTasks, err := todoist.GetTasks(tdistClient)
	if err != nil {
		panic("tasks from todoist not found")
	}

	for _, v := range *tdTasks {
		fmt.Println(v.Content, v.Id, v.CreatedAt)
	}
}
