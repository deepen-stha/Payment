package rabbitmq

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"payment/model"
	"payment/rediscache"
	"payment/repository"

	//importing payment/redis as redis_cache

	"bitbucket.org/kaleyra/mongo-sdk/mongo"
	"github.com/go-redis/redis"
	"github.com/streadway/amqp"
)

//consuming the message it accepts the conn adn the redisClient
//so that we can fetch the required data from the redis
func ConsumeMessage(conn *amqp.Connection, redisClient *redis.Client, sqlconn *sql.DB, collection *mongo.Collection) {
	//creating a channel
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	//consuming the message that are in the queue
	msgs, err := ch.Consume(
		"PaymentQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	var details model.Details
	// By default, sends and receives block until the other side is ready. This allows goroutines to synchronize without explicit locks or condition variables.
	channel := make(chan bool)
	//creating a go routine
	go func() {
		for d := range msgs {
			fmt.Printf("Received Message: %s\n", d.Body)
			err2 := json.Unmarshal(d.Body, &details)

			if err2 != nil {
				fmt.Println("error occurs while unmarshalling the data", err)
				// log.Fatal(err2)
			}
			//calling redis cache to get the data
			data := rediscache.GetCacheData(redisClient, details)
			//if data is present in the cache
			if data != false {
				fmt.Println("data is present there")
				//calling the redis cache to set the data
				// rediscache.SetRedisValue(redisClient, details)
				repository.AddMongoData(collection, details)
				continue
			}
			//if data is not in the cache then look mysql database for data
			repository.GetDataByAcc(redisClient, sqlconn, details)
			repository.AddMongoData(collection, details)
		}
	}()
	<-channel
}
