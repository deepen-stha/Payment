package main

import (
	"database/sql"
	"fmt"
	"payment/rabbitmq"
	"payment/repository"
)

var SQL_CONN *sql.DB

func main() {

	initEnv()

	initLogger()

	SQL_CONN := InitMySQL()
	repository.GetSQlData(SQL_CONN)

	//function to initialize the mongodb-sdk
	coll := initMongoSDK()
	fmt.Println(coll)

	//function to initialize the rabbitmq
	rmqConn := initRMQ()
	defer rmqConn.Close()

	//function to initialize the redis
	redisClient := initRedis()
	fmt.Println(redisClient)

	//pushing msg into the RMQ
	rabbitmq.PublishMessage(rmqConn)
	//consuming the msg
	rabbitmq.ConsumeMessage(rmqConn, redisClient,SQL_CONN,coll)

}
