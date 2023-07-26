package tools

import (
	"fmt"
	"github.com/spf13/viper"
)

func AsRef[T any](a T) *T { return &a }

func GetStringFromEnv(key string) (string, error) {
	if !viper.IsSet(key) {
		return "", fmt.Errorf("environment variable %s not found", key)
	}

	value := viper.Get(key)
	return fmt.Sprintf("%v", value), nil
}
