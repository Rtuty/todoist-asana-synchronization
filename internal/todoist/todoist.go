package todoist

import (
	"fmt"

	"github.com/volyanyk/todoist"
)

// NewTdIstClient возвращает подключение(клиент) к api todoist
func NewTdIstClient() (*todoist.Client, error) {
	token, err := initTodoistToken()
	if err != nil {
		return nil, err
	}

	client := todoist.New(token)

	return client, nil
}

// GetTasks
func GetTasks(client *todoist.Client) (*[]todoist.Task, error) {
	// Получаем список проектов
	projects, err := client.GetProjects()
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении списка проектов. details: %v", err)
	}

	//  Находим id проекта (папки) по наименованию, которое указано в .env
	folderName, err := initFolderName()
	if err != nil {
		return nil, err
	}

	var folderId string

	for _, v := range *projects {
		if v.Name == folderName {
			folderId = v.ID
		}
	}

	// Получаем список задач из указанного проекта
	folderTasks, err := client.GetActiveTasks(todoist.GetActiveTasksRequest{ProjectId: folderId})
	if err != nil {
		return nil, fmt.Errorf("не удалось получить задачи из указанного проекта. details: %v", err)
	}

	return folderTasks, nil
}
