package app

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"todoistapi/internal/db"
	asn "todoistapi/internal/task-managers/asana"
	"todoistapi/internal/task-managers/todoist"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
}

func RunInstance() {
	ctx := context.Background()

	vpr := viper.New()
	vpr.AutomaticEnv()

	rc := db.RedisCli{
		Addr: vpr.GetString("REDIS_PASS"),
		Pass: vpr.GetString("REDIS_ADDR"),
		V:    vpr,
	}

	as := asn.GetAsanaFnc(vpr)
	td := todoist.GetTodoistFnc(vpr)

	rdb, err := rc.NewRedisClient()
	if err != nil {
		log.Fatalf("get redis client error: %v", err)
	}

	h := CmdHandler{
		Ctx:   ctx,
		Vpr:   vpr,
		Redis: rdb,
		TdsFn: td,
		AsnFn: as,
	}

	h.CommandLineHandler()
}
