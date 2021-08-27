package repository

import (
	"fmt"
	"payment/model"

	"bitbucket.org/kaleyra/mongo-sdk/mongo"
)

func AddMongoData(collection *mongo.Collection, data model.Details) {

	insertId, err := collection.InsertOne(data)
	if err != nil {
		fmt.Println("error while inserting data : ", err)
	}
	fmt.Println("after inserting:", insertId)
}
