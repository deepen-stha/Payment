package rediscache

import (
	"fmt"
	"payment/model"

	"github.com/go-redis/redis"
)

func GetCacheData(client *redis.Client, data model.Details) bool {

	val, err := client.Get(string(data.AccountNo)).Result()
	if err != nil {
		fmt.Println(err)
		return false
	}
	fmt.Println("value is found: ", val)
	return true
}
