package db

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

type RedisCli struct {
	V *viper.Viper
}

// NewRedisClient возвращает подключение(клиент) к базе данных redis
func (rdc *RedisCli) NewRedisClient() (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     rdc.V.GetString("REDIS_ADDR"),
		Password: rdc.V.GetString("REDIS_PASS"),
		DB:       0, //todo
	})

	return rdb, nil
}
