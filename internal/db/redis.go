package db

import (
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

type RedisCli struct {
	Addr string
	Pass string
	V    *viper.Viper
}

// NewRedisClient возвращает подключение(клиент) к базе данных redis
func (rdc *RedisCli) NewRedisClient() (*redis.Client, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     rdc.Addr,
		Password: rdc.Pass,
		DB:       0, //todo
	})

	return rdb, nil
}
