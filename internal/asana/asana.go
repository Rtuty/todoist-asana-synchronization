package asana

import (
	"errors"
	"fmt"
	"todoistapi/tools"

	"bitbucket.org/mikehouston/asana-go"
)

// Получаем токен, создаем и возвращаем новый клиент
func NewClient() (*asana.Client, error) {
	token, err := initAsanaToken()
	if err != nil {
		return nil, err
	}

	client := asana.NewClientWithAccessToken(token)

	return client, nil
}

// Получаем нужный workspace
func GetWorkSpace(client *asana.Client) (string, error) {
	workSpaceName, err := initAsanaWorkSpace()
	if err != nil {
		return "", err
	}

	mass, err := client.AllWorkspaces(&asana.Options{Pretty: tools.AsRef(true)})
	if err != nil {
		return "", err
	}

	for _, v := range mass {
		if v.Name == workSpaceName {
			return v.ID, nil
		}
	}

	return "", errors.New("work space not found")
}

// Получить id пользователя asana по имени из .env
func GetUserIdByName(client *asana.Client) (string, error) {
	mass, err := client.AllWorkspaces(&asana.Options{Pretty: tools.AsRef(true)})
	if err != nil {
		return "", err
	}

	userName, err := initUserName()
	if err != nil {
		return "", err
	}

	for _, v := range mass {
		u, err := v.AllUsers(client)
		if err != nil {
			return "", err
		}

		for _, v2 := range u {
			if v2.Name == userName {
				return v2.ID, nil
			}
		}
	}

	return "", errors.New("get user id not found")
}

// todo: return and error handeling
func GetTasksByUserId(client *asana.Client, userId string) {
	workspace, err := GetWorkSpace(client)
	if err != nil {
		panic(err)
	}

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
