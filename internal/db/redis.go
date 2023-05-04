package db

import (
	"context"
	"errors"
	"fmt"
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

func NewClient() (*redis.Client, error) {
	addr, passw, err := initAddrPass()
	if err != nil {
		return nil, err
	}

	rdb := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: passw,
		DB:       0,
	})

	return rdb, nil
}

// Тестовая функция для проверки корректности работы БД
func GetRedisClient(rdb *redis.Client) {
	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}