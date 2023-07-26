package app

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"strconv"
	"todoistapi/internal/db"
	"todoistapi/internal/task-managers/asana"
	"todoistapi/internal/task-managers/todoist"
	"todoistapi/internal/tools"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func RunInstance() {
	ctx := context.Background()
	viperCon := viper.New()
	viperCon.AutomaticEnv()

	rc := db.RedisCli{V: viperCon}
	td := todoist.GetTodoistFnc(viperCon)

	rdb, err := rc.NewRedisClient()
	if err != nil {
		log.Fatalf("get redis client error: %v", err)
	}

	tdCl, err := td.NewTdIstClient()
	if err != nil {
		log.Fatalf("get todoist client error: %v", err)
	}

	asCl, err := asn.NewAsanaClient()
	if err != nil {
		log.Fatalf("get asana client error: %v", err)
	}

	folderName, err := tools.GetStringFromEnv("FOLDERNAME")
	if err != nil {
		log.Fatal(err)
	}

	tdTasks, err := td.GetTasksByFolderId(tdCl, todoist.Args{FolderName: folderName, Completed: true})
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
