package asn

import (
	"bitbucket.org/mikehouston/asana-go"
	"github.com/spf13/viper"
)

type AsanaFnc interface {
	NewAsanaClient() (*asana.Client, error)
	GetWorkSpace(client *asana.Client) (string, error)
	GetUserIdByName(client *asana.Client) (string, error)
	GetUncompletedTasks(client *asana.Client, userId string, workSpaceId string) ([]asana.Task, error)
}

type AsanaCli struct {
	V *viper.Viper
}

func GetAsanaFnc(v *viper.Viper) AsanaFnc {
	return &AsanaCli{V: v}
}
