package todoist

import (
	"fmt"
	"todoistapi/internal/tools"

	"github.com/volyanyk/todoist"
)

// NewTdIstClient возвращает подключение(клиент) к api todoist
func (td *TodoistCli) NewTdIstClient() (*todoist.Client, error) {
	token, err := tools.GetStringFromEnv("TOKENTODOIST")
	if err != nil {
		return nil, err
	}

	return todoist.New(token), nil
}

// GetTasksByFolderId возвращает список проектов
func (td *TodoistCli) GetTasksByFolderId(client *todoist.Client, args Args) (*[]todoist.Task, error) {
	projects, err := client.GetProjects() // Получаем список проектов
	if err != nil {
		return nil, fmt.Errorf("error when getting the list of projects. details: %v", err)
	}

	var folderId string

	for _, v := range *projects { //  Находим id проекта (папки) по наименованию, которое указано в .env
		if v.Name == args.FolderName {
			folderId = v.ID
		}
	}

	// Получаем список задач из указанного проекта
	folderTasks, err := client.GetActiveTasks(todoist.GetActiveTasksRequest{ProjectId: folderId})
	if err != nil {
		return nil, fmt.Errorf("failed to get tasks from the specified project. details: %v", err)
	}

	return folderTasks, nil
}
