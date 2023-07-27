package app

import (
	"bitbucket.org/mikehouston/asana-go"
	"bufio"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"github.com/volyanyk/todoist"
	"log"
	"os"
	"strings"
	asn "todoistapi/internal/task-managers/asana"
	tdist "todoistapi/internal/task-managers/todoist"
)

type CmdHandler struct {
	Ctx   context.Context
	Vpr   *viper.Viper
	Redis *redis.Client

	TdsFn tdist.TodoistFnc
	AsnFn asn.AsanaFnc
}

func (h *CmdHandler) CommandLineHandler() {
	scn := bufio.NewScanner(os.Stdin)

	asCl, err := h.AsnFn.NewAsanaClient()
	if err != nil {
		log.Fatalf("get asana client error: %v", err)
	}

	tdCl, err := h.TdsFn.NewTdIstClient()
	if err != nil {
		log.Fatalf("get todoist client error: %v", err)
	}

	for {
		fmt.Print("command: \n")
		scn.Scan()

		cmd := scn.Text()

		cmd = strings.TrimSuffix(cmd, "\n")

		switch cmd {
		case "sync_td":
			go h.asanaToTodoistSyncTasks(h.Ctx, asCl, tdCl)
		case "exit":
			fmt.Println("program break")
			return
		}
	}
}

// asanaToTodoistSyncTasks Переносит незавершенные задачи из проекта asana в todoist
func (h *CmdHandler) asanaToTodoistSyncTasks(ctx context.Context, aCl *asana.Client, tCl *todoist.Client) {
	usrId, err := h.AsnFn.GetUserIdByName(aCl)
	if err != nil {
		log.Print(err)
	}

	wrkSpName, err := h.AsnFn.GetWorkSpace(aCl)
	if err != nil {
		log.Print(err)
	}

	asnTsks, err := h.AsnFn.GetUncompletedTasks(aCl, usrId, wrkSpName)
	if err != nil {
		log.Print(err)
	}

	for _, t := range asnTsks {
		new := todoist.AddTaskRequest{
			Content:     t.Name,
			Description: "Asana id:" + t.ID,
		}

		_, err := tCl.AddTaskContext(new, ctx)
		if err != nil {
			log.Print(err)
		}
	}
}
