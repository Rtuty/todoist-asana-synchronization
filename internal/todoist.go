package internal

import (
	"fmt"
	"os"

	"github.com/volyanyk/todoist"
)

func TodoistFunctionality() {
	// Подключаемся к API TODOIST'a
	apiTodoist, exists := os.LookupEnv("TOKENTODOIST")
	if !exists {
		panic("Todoist API не найден в .env")
	}

	api := todoist.New(apiTodoist)

	// Получаем список проектов
	projects, err := api.GetProjects()
	if err != nil {
		panic("Ошибка при получении списка проектов")
	}

	//  Находим id по наименованию, которое указано в .env
	folderName, exists := os.LookupEnv("FOLDERNAME")
	if !exists {
		panic("Наименование рабочей папки не найдено в файле .env")
	}

	var jobId string

	for _, v := range *projects {
		if v.Name == folderName {
			jobId = v.ID
		}
	}

	// Получаем список задач из указанного проекта
	jobTasks, err := api.GetActiveTasks(todoist.GetActiveTasksRequest{ProjectId: jobId})
	if err != nil {
		panic("Не удалось получить задачи из указанного проекта")
	}

	fmt.Printf("%v \n", jobTasks)
}
