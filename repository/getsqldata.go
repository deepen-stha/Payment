package repository

import (
	"database/sql"
	"fmt"
	"log"
	"payment/model"
	"payment/rediscache"

	"github.com/go-redis/redis"
)

func GetSQlData(db *sql.DB) {

	res, err := db.Query("SELECT * FROM DETAIL")

	defer res.Close()

	if err != nil {
		fmt.Println(err)
	}

	for res.Next() {

		var city model.Details
		err := res.Scan(&city.AccountNo, &city.Name, &city.Pay)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v\n", city)
	}
}

func GetDataByAcc(redisClient *redis.Client, db *sql.DB, data model.Details) {
	//QUERY TO GET THE DATA ON THE BASIS OF ACCOUNT NO
	res, err := db.Query("SELECT * FROM DETAIL WHERE ACC_NO = ?", data.AccountNo)
	defer res.Close()
	if err != nil {
		fmt.Println("error while running query")
	}
	//it will give only one id
	for res.Next() {

		var dbdata model.Details
		err := res.Scan(&dbdata.AccountNo, &dbdata.Name, &dbdata.Pay)

		if err != nil {
			log.Fatal(err)
		}

		//checking if the available balance is greater than the balance to spend
		if dbdata.Pay > data.Pay {
			var newamount = dbdata.Pay - data.Pay
			insForm, err := db.Prepare("UPDATE DETAIL SET BALANCE=? WHERE ACC_NO=?")
			if err != nil {
				panic(err.Error())
			}
			insForm.Exec(newamount, data.AccountNo)
			log.Println("UPDATE: Amount: ", newamount)
			rediscache.SetRedisValue(redisClient, data)
			return
		}
		fmt.Println("The amount exceeds the available balance.")
	}
}
