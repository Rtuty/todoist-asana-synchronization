package asana

import (
	"fmt"
	"os"
	"todoistapi/tools"

	"bitbucket.org/mikehouston/asana-go"
)

// Получаем токен, создаем и возвращаем новый клиент
func NewClient() *asana.Client {
	token, exists := os.LookupEnv("ASANA_TOKEN")
	if !exists {
		panic("Asana API token not found")
	}

	client := asana.NewClientWithAccessToken(token)

	return client
}

// Получаем нужный workspace
func GetWorkSpace(client *asana.Client) string {
	workSpaceName, exists := os.LookupEnv("WORKSPACE_NAME")
	if !exists {
		panic("workSpaceName doesn't exist")
	}

	mass, err := client.AllWorkspaces(&asana.Options{Pretty: tools.AsRef(true)})
	if err != nil {
		panic(err)
	}

	for _, v := range mass {
		if v.Name == workSpaceName {
			return v.ID
		}
	}
	panic("GetWorkSpace error")
}

// Получить id пользователя asana по имени из .env
func GetUserIdByName(client *asana.Client) (string, error) {
	mass, err := client.AllWorkspaces(&asana.Options{Pretty: tools.AsRef(true)})
	if err != nil {
		return "", err
	}

	var userId string

	userName, exists := os.LookupEnv("USER_NAME")
	if !exists {
		return "", err
	}

	for _, v := range mass {
		u, err := v.AllUsers(client)
		if err != nil {
			return "", err
		}

		for _, v2 := range u {
			if v2.Name == userName {
				userId = v2.ID
				break
			}
		}
	}

	return userId, nil
}

func GetTasksByUserId(client *asana.Client, userId string) {
	var workspace string = GetWorkSpace(client)

	tasks, _, err := client.QueryTasks(
		&asana.TaskQuery{
			Workspace: workspace,
			Assignee:  userId,
		},
		&asana.Options{
			Pretty: tools.AsRef(true),
		},
	)
	if err != nil {
		panic(err)
	}

	//todo: не возвращает completed, completedat. Возможно, придется дописывать пакет и отправлять pull request
	for _, v := range tasks {
		fmt.Println(v.Name, v.Completed, v.CompletedAt, v.Projects, v.ResourceSubtype)
	}
}

func TestFunc(client *asana.Client, userId string) {
	p := &asana.Project{}
	tasks, _, err := p.Tasks(client, &asana.Options{Debug: tools.AsRef(true)})
	if err != nil {
		panic(err)
	}

	fmt.Println(tasks)
}
