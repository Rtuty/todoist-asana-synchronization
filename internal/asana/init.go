package asana

import (
	"errors"
	"os"
)

func initAsanaToken() (string, error) {
	token, exists := os.LookupEnv("ASANA_TOKEN")
	if !exists {
		return "", errors.New("asana API token not found in .env")
	}

	return token, nil
}

func initAsanaWorkSpace() (string, error) {
	workSpaceName, exists := os.LookupEnv("WORKSPACE_NAME")
	if !exists {
		return "", errors.New("asana workspace name not found in .env")
	}

	return workSpaceName, nil
}

func initUserName() (string, error) {
	userName, exists := os.LookupEnv("USER_NAME")
	if !exists {
		return "", errors.New("asana user name not found in .env")
	}

	return userName, nil
}
