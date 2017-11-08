package redisDB

import (
	"fmt"
	"github.com/go-redis/redis"
	"time"
)

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Connect Redis Successify!")
}

func Set(key string, value interface{}, duration time.Duration) {
	err := client.Set(key, value, duration).Err()
	if err != nil {
		panic(err)
	}
}

func Get(key string) interface{} {
	val, err := client.Get(key).Result()
	if err == redis.Nil {
		return err
	} else if err != nil {
		panic(err)
	}
	return val
}
