package app

import (
	"context"
	"github.com/joho/godotenv"
	"log"
	"strconv"
	"todoistapi/internal/asana"
	"todoistapi/internal/db"
	"todoistapi/internal/todoist"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func RunInstance() {
	ctx := context.Background()

	rdb, err := db.NewRedisClient()
	if err != nil {
		log.Fatalf("get redis client error: %v", err)
	}

	tdCl, err := todoist.NewTdIstClient()
	if err != nil {
		log.Fatalf("get todoist client error: %v", err)
	}

	asCl, err := asn.NewAsanaClient()
	if err != nil {
		log.Fatalf("get asana client error: %v", err)
	}

	tdTasks, err := todoist.GetTasks(tdCl)
	if err != nil {
		log.Fatalf("todoist get tasts error: %v", err)
	}

	for _, v := range *tdTasks {
		log.Println(v.Content, v.Id, v.CreatedAt)
	}

	usName, err := asn.GetUserIdByName(asCl)
	if err != nil {
		panic(err)
	}

	wsId, err := asn.GetWorkSpace(asCl)
	if err != nil {
		log.Printf("get workspace error: %v", err)
	}

	tasks, err := asn.GetUncompletedTasks(asCl, usName, wsId)
	if err != nil {
		log.Printf("get uncompleted tasks error: %v", err)
	}

	for _, v := range tasks {
		if err := rdb.Set(
			ctx,
			v.ID,
			[]string{
				v.ID,
				v.Name,
				strconv.FormatBool(*v.Completed),
				v.Parent.ID,
			}, 0).Err(); err != nil {
			log.Printf("add redis error:%v, id: %s", err, v.ID)
		}
	}
}
