package redis

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

func Connect() {
	fmt.Println("get redis client")

	redisAddr, exists := os.LookupEnv("REDIS_ADDR")
	if !exists {
		panic("redis options address doesn't exist")
	}

	redisPassword, exists := os.LookupEnv("REDIS_PASS")
	if !exists {
		panic("redis options password doesn't exist")
	}

	client := redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: redisPassword,
		DB:       0,
	})

	pong, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}

	fmt.Println(pong)
}
