package main

import (
	"log"
	"todoistapi/internal/asana"
	"todoistapi/internal/db"
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
	rdb, err := db.NewClient()
	if err != nil {
		log.Fatalf("get redis client error: %v", err)
	}

	if err = db.GetRedisClient(rdb); err != nil {
		log.Fatalf("redis test function error: %v", err)
	}

	tdistClient, err := todoist.NewClient()
	if err != nil {
		log.Fatalf("todoist get client error: %v", err)
	}

	tdTasks, err := todoist.GetTasks(tdistClient)
	if err != nil {
		log.Fatalf("todoist get tasts error: %v", err)
	}

	for _, v := range *tdTasks {
		log.Println(v.Content, v.Id, v.CreatedAt)
	}

	aCl, err := asana.NewClient()
	if err != nil {
		panic(err)
	}

	usName, err := asana.GetUserIdByName(aCl)
	if err != nil {
		panic(err)
	}

	asana.GetTasksByUserId(aCl, usName)
}
