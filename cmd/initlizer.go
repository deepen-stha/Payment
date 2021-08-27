package main

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"payment/constant"

	"github.com/go-redis/redis"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
	"go.uber.org/zap"

	// "bitbucket.org/kaleyra/mongo-sdk/bson"
	"bitbucket.org/kaleyra/mongo-sdk/mongo"
)

var logger *zap.Logger

//function to initialize the .env file
func initEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Printf("Failed to load the env file : Error = %s", err.Error())
		return
	}
}

//function to initialize the logger
func initLogger() {
	logger = zap.NewExample()
	defer logger.Sync()
	logger.Info("failed to fetch URL",
		zap.String("url", "http://example.com"),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}

// funciton to initialize the mysql database
func InitMySQL() *sql.DB {
	db, err := sql.Open("mysql", "root:"+constant.SQL_PASS+"@tcp(127.0.0.1:3306)/customers")
	// defer db.Close()

	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	fmt.Println("successfully connected to mysql")
	return db
}

//function to initialize the mongosdk
func initMongoSDK() *mongo.Collection {
	db := mongo.URI{
		Username: "",
		Password: "",
		Host:     os.Getenv("DB_HOST"),
		DB:       os.Getenv("DB_NAME"),
		Port:     os.Getenv("DB_PORT"),
	}
	client, err := mongo.NewClient(db)
	if err != nil {
		fmt.Println("not able to connect to mongodb-sdk")
		// sugarLogger.Errorf("Failed to connect to mongodb = %s", err.Error())
		return nil
	}
	fmt.Println("Connected to MongoDB-SDK!")
	collection := client.Collection(os.Getenv("DB_COLLECTION"))
	return collection
}

//function to initialize the redis
func initRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     constant.CLOUD_ADDR,
		Password: constant.CLOUD_PASS,
		DB:       0,
	})
	return client
}

//function to initialize the RMQ
func initRMQ() *amqp.Connection {

	//tryinng to connect to our rabbitmq
	conn, err := amqp.Dial("amqp://" + os.Getenv("USER_NAME") + ":" + os.Getenv("PASSWORD") + "@10.20.30.25:5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
		// return nil
	}
	fmt.Println("Successfully connected to our RabbitMq")
	return conn

}
