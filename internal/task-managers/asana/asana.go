package asn

import (
	"bitbucket.org/mikehouston/asana-go"
	"errors"
	"fmt"
	"todoistapi/internal/tools"
)

// NewAsanaClient возврщает токен, создаем и возвращаем новый клиент
func NewAsanaClient() (*asana.Client, error) {
	token, err := initAsanaToken()
	if err != nil {
		return nil, err
	}

	return asana.NewClientWithAccessToken(token), nil
}

// GetWorkSpace возвращает id workspace по имени из .env
func GetWorkSpace(client *asana.Client) (string, error) {
	workSpaceName, err := initAsanaWorkSpace()
	if err != nil {
		return "", fmt.Errorf("init workspace name error in get workspace function: %v", err)
	}

	mass, err := client.AllWorkspaces(&asana.Options{Pretty: tools.AsRef(true)})
	if err != nil {
		return "", fmt.Errorf("get workspaces error: %v", err)
	}

	for _, v := range mass {
		if v.Name == workSpaceName {
			return v.ID, nil
		}
	}

	return "", errors.New("work space not found")
}

// GetUserIdByName получает id пользователя asana по имени из .env
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

// GetUncompletedTasks возвращает незакрытые задачи по userId и workspace id
func GetUncompletedTasks(client *asana.Client, userId string, workSpaceId string) ([]asana.Task, error) {
	tasks, _, err := client.QueryTasks(
		&asana.TaskQuery{
			Workspace: workSpaceId,
			Assignee:  userId,
		},

		asana.Fields(asana.Task{}),

		&asana.Options{
			Pretty: tools.AsRef(true),
		},
	)
	if err != nil {
		return []asana.Task{}, err
	}

	var res []asana.Task

	for _, v := range tasks {
		if !*v.Completed {
			res = append(res, *v)
		}
	}

	return res, nil
}
