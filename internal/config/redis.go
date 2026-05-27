package config

import(
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

var Ctx = context.Background()


func ConnectRedis() *redis.Client{

	rdb := redis.NewClient(&redis.Options{
       
		Addr: "127.0.0.1:6378",
	})

	_, err := rdb.Ping(Ctx).Result()

	if err != nil {
		panic(err)

	}

	fmt.Println("Connected to Redis...")

	return rdb
}