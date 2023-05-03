package main

import (
	"fmt"
	"log"
	redisdb "todoistapi/internal/db/redis"

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
	fmt.Println("----------------------redis test----------------------------")

	rdb, err := redisdb.NewClient()
	if err != nil {
		panic(err)
	}

	redisdb.GetRedisClient(rdb)

	fmt.Println("----------------------redis test----------------------------")

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
