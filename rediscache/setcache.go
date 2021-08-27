package rediscache

import (
	"fmt"
	"payment/model"
	"time"

	"github.com/go-redis/redis"
)

func SetRedisValue(client *redis.Client, data model.Details) {

	fmt.Println(client)
	err := client.Set(string(data.AccountNo), data, time.Second)
	if err != nil {
		fmt.Println(err)
		return
	}
}
