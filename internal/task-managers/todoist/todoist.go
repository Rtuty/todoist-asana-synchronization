package todoist

import (
	"github.com/spf13/viper"
	"github.com/volyanyk/todoist"
)

type TodoistFnc interface {
	NewTdIstClient() (*todoist.Client, error)
	GetTasksByFolderId(client *todoist.Client, args Args) (*[]todoist.Task, error)
}

type TodoistCli struct {
	V *viper.Viper
}

func GetTodoistFnc(v *viper.Viper) TodoistFnc {
	return &TodoistCli{
		V: v,
	}
}

type Args struct {
	FolderName string
	Completed  bool
}
