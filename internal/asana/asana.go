package asana

import (
	"os"

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

// Получить id пользователя asana по имени из .env
func GetUserIdByName(client *asana.Client) (string, error) {
	var pretty bool = true

	mass, err := client.AllWorkspaces(&asana.Options{Pretty: &pretty})
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
