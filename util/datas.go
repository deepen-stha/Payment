package util

import (
	"payment/model"
)

//getting all the msg and returning it
func GetMessages() []model.Details {

	//forming the json messages that we want to publish
	users := []model.Details{
		{
			AccountNo: 1,
			Name:      "Deepen Shrestha",
			Pay:       1000,
		},
		{
			AccountNo: 2,
			Name:      "Dipesh",
			Pay:       2000,
		},
		{
			AccountNo: 3,
			Name:      "David",
			Pay:       3000,
		},
		{
			AccountNo: 4,
			Name:      "Viram",
			Pay:       4000,
		},
	}

	// fmt.Println("length of users are: ", len(users))
	return users
}
