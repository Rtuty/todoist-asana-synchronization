package app

import (
	"context"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
	"log"
	"todoistapi/internal/db"
	asn "todoistapi/internal/task-managers/asana"
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

	vpr := viper.New()
	vpr.AutomaticEnv()

	rAddr, err := tools.GetStringFromEnv("REDIS_ADDR")
	if err != nil {
		log.Fatal(err)
	}

	rPass, err := tools.GetStringFromEnv("REDIS_PASS")
	if err != nil {
		log.Fatal(err)
	}

	rc := db.RedisCli{
		Addr: rAddr,
		Pass: rPass,
		V:    vpr,
	}

	rdb, err := rc.NewRedisClient()
	if err != nil {
		log.Fatalf("get redis client error: %v", err)
	}

	as := asn.GetAsanaFnc(vpr)
	td := todoist.GetTodoistFnc(vpr)

	h := CmdHandler{
		Ctx:   ctx,
		Vpr:   vpr,
		Redis: rdb,
		TdsFn: td,
		AsnFn: as,
	}

	h.CommandLineHandler()
}
