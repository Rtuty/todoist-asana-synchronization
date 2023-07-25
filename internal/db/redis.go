package db

import (
	"context"
	"errors"
	"os"

	"github.com/redis/go-redis/v9"
)

// Иннициализация данных из .env
var ctx = context.Background()

func initAddrPass() (string, string, error) {
	addr, exists := os.LookupEnv("REDIS_ADDR")
	if !exists {
		return "", "", errors.New("не удалось получить адрес для redis из файла .env")
	}

	passw, exists := os.LookupEnv("REDIS_PASS")
	if !exists {
		return "", "", errors.New("не удалось получить пароль для redis из файла .env")
	}

	return addr, passw, nil
}

// NewRedisClient возвращает подключение(клиент) к базе данных redis
func NewRedisClient() (*redis.Client, error) {
	addr, passw, err := initAddrPass()
	if err != nil {
		return nil, err
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: passw,
		DB:       0, //todo
	})

	return rdb, nil
}
