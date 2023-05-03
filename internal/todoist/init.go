package todoist

import (
	"errors"
	"os"
)

func initTodoistToken() (string, error) {
	token, exists := os.LookupEnv("TOKENTODOIST")
	if !exists {
		return "", errors.New("todoist API token not found in .env")
	}

	return token, nil
}

func initFolderName() (string, error) {
	folderName, exists := os.LookupEnv("FOLDERNAME")
	if !exists {
		return "", errors.New("folder name not found in .env")
	}

	return folderName, nil
}
